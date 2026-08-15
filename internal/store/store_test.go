package store

import (
	"testing"

	"metrics/internal/model"
)

func TestRecordGet(t *testing.T) {
	s := New()
	if err := s.Record(&model.Sample{ID: "a", Name: "cpu"}); err != nil {
		t.Fatal(err)
	}
	if err := s.Record(&model.Sample{ID: "a"}); err != ErrSampleExists {
		t.Fatalf("dup err=%v", err)
	}
	if _, err := s.Get("a"); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Get("nope"); err != ErrSampleNotFound {
		t.Fatalf("err=%v", err)
	}
	if s.Count() != 1 {
		t.Fatalf("Count=%d", s.Count())
	}
}

func TestOrderIDsFresh(t *testing.T) {
	s := New()
	_ = s.Record(&model.Sample{ID: "b", Name: "x"})
	_ = s.Record(&model.Sample{ID: "a", Name: "x"})
	ids := s.OrderIDs()
	ids[0] = "z"
	if s.OrderIDs()[0] != "b" {
		t.Fatal("OrderIDs aliased")
	}
}
