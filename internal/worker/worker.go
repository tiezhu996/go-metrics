package worker

import (
	"context"
	"sync"

	"metrics/internal/model"
	"metrics/internal/service"
)

type Aggregator interface {
	Aggregate(ctx context.Context, sm *model.Sample) (model.Summary, error)
}

type Pool struct {
	svc     *service.Service
	agg     Aggregator
	workers int
}

func New(svc *service.Service, a Aggregator, workers int) *Pool {
	if workers <= 0 {
		workers = 1
	}
	return &Pool{svc: svc, agg: a, workers: workers}
}

func (p *Pool) Run(ctx context.Context) model.Summary {
	buckets := p.svc.ListBuckets()

	var wg sync.WaitGroup
	ch := make(chan []*model.Sample, len(buckets))

	go func() {
		defer close(ch)
		for _, b := range buckets {
			select {
			case <-ctx.Done():
				return
			case ch <- b:
			}
		}
	}()

	results := make(chan model.Summary, p.workers)

	for i := 0; i < p.workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var local model.Summary
			for bucket := range ch {
				for _, sm := range bucket {
					part, err := p.agg.Aggregate(ctx, sm)
					if err != nil {
						local.Failed++
						continue
					}
					local = model.MergeSummary(local, part)
				}
			}
			results <- local
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var sum model.Summary
	for local := range results {
		sum = model.MergeSummary(sum, local)
	}
	return sum
}
