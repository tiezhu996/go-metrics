package worker

import (
	"context"
	"fmt"
	"testing"

	"metrics/internal/config"
	"metrics/internal/model"
	"metrics/internal/service"
	"metrics/internal/store"
)

type okAgg struct{}

func (okAgg) Aggregate(ctx context.Context, sm *model.Sample) (model.Summary, error) {
	if err := ctx.Err(); err != nil {
		return model.Summary{}, err
	}
	return model.Summary{Count: 1, Sum: sm.Value, Max: sm.Value, Min: sm.Value}, nil
}

func newPool() (*service.Service, *Pool) {
	s := store.New()
	svc := service.New(s, config.Load())
	return svc, New(svc, okAgg{}, 4)
}

func TestRunSummary(t *testing.T) {
	svc, p := newPool()
	for i := 0; i < 10; i++ {
		_ = svc.Record(&model.Sample{ID: fmt.Sprintf("s%d", i), Name: "cpu", Value: float64(i)})
	}
	sum := p.Run(context.Background())
	if sum.Count != 10 || sum.Sum != 45 {
		t.Fatalf("sum=%+v", sum)
	}
}

type failAgg struct{ failID string }

func (f failAgg) Aggregate(ctx context.Context, sm *model.Sample) (model.Summary, error) {
	if sm.ID == f.failID {
		return model.Summary{}, fmt.Errorf("fail %s", sm.ID)
	}
	return model.Summary{Count: 1, Sum: sm.Value, Max: sm.Value, Min: sm.Value}, nil
}

func TestRunFailed(t *testing.T) {
	svc, _ := newPool()
	_ = svc.Record(&model.Sample{ID: "ok", Name: "cpu", Value: 1})
	_ = svc.Record(&model.Sample{ID: "bad", Name: "cpu", Value: 2})
	p := New(svc, failAgg{"bad"}, 2)
	sum := p.Run(context.Background())
	if sum.Failed != 1 || sum.Count != 1 {
		t.Fatalf("sum=%+v", sum)
	}
}

func TestRunCancel(t *testing.T) {
	svc, p := newPool()
	_ = svc.Record(&model.Sample{ID: "s0", Name: "cpu", Value: 1})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	sum := p.Run(ctx)
	if sum.Count != 0 {
		t.Fatalf("Count=%d want 0", sum.Count)
	}
}
