package service

import (
	"errors"
	"fmt"

	"metrics/internal/config"
	"metrics/internal/model"
	"metrics/internal/store"
)

type Service struct {
	store      *store.Store
	bucketSize int
}

func New(s *store.Store, cfg *config.Config) *Service {
	b := cfg.BucketSize
	if b <= 0 {
		b = 1
	}
	return &Service{store: s, bucketSize: b}
}

func (svc *Service) Record(sm *model.Sample) error {
	if !model.ValidSample(sm) {
		return errors.New("invalid sample")
	}
	if err := svc.store.Record(sm); err != nil {
		return fmt.Errorf("record %s: %v", sm.ID, err)
	}
	return nil
}

func (svc *Service) ListBuckets() [][]*model.Sample {
	ss := svc.store.ListSamples()
	model.SortSamples(ss)
	return model.BuildBuckets(ss, svc.bucketSize)
}

func (svc *Service) Summary() (model.Summary, error) {
	ss := svc.store.ListSamples()
	if len(ss) == 0 {
		return model.Summary{}, errors.New("no samples")
	}
	var sum model.Summary
	for _, sm := range ss {
		part := model.Summary{Count: 1, Sum: sm.Value, Max: sm.Value, Min: sm.Value}
		sum = model.MergeSummary(sum, part)
	}
	return sum, nil
}
