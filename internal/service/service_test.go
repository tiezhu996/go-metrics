package service

import (
	"errors"
	"testing"

	"metrics/internal/config"
	"metrics/internal/model"
	"metrics/internal/store"
)

func newSvc() (*store.Store, *Service) {
	s := store.New()
	return s, New(s, config.Load())
}

func TestRecordSummary(t *testing.T) {
	_, svc := newSvc()
	_ = svc.Record(&model.Sample{ID: "a", Name: "cpu", Value: 1})
	_ = svc.Record(&model.Sample{ID: "b", Name: "cpu", Value: 3})
	if err := svc.Record(&model.Sample{}); err == nil {
		t.Fatal("invalid sample should error")
	}
	sum, err := svc.Summary()
	if err != nil || sum.Count != 2 || sum.Sum != 4 || sum.Max != 3 || sum.Min != 1 {
		t.Fatalf("sum=%+v err=%v", sum, err)
	}
}

func TestListBucketsOrder(t *testing.T) {
	_, svc := newSvc()
	_ = svc.Record(&model.Sample{ID: "b", Name: "cpu", Value: 1})
	_ = svc.Record(&model.Sample{ID: "a", Name: "cpu", Value: 1})
	buckets := svc.ListBuckets()
	if len(buckets) == 0 || buckets[0][0].ID != "a" {
		t.Fatalf("buckets=%v", buckets)
	}
}

func TestRecordWraps(t *testing.T) {
	_, svc := newSvc()
	_ = svc.Record(&model.Sample{ID: "a", Name: "cpu"})
	if err := svc.Record(&model.Sample{ID: "a", Name: "cpu"}); !errors.Is(err, store.ErrSampleExists) {
		t.Fatalf("errors.Is=false err=%v", err)
	}
}
