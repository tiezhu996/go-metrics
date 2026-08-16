package store

import (
	"errors"
	"sync"

	"metrics/internal/model"
)

var (
	ErrSampleNotFound = errors.New("sample not found")
	ErrSampleExists   = errors.New("sample already exists")
)

type Store struct {
	mu      sync.RWMutex
	samples map[string]*model.Sample
	order   []string
}

func New() *Store {
	return &Store{
		samples: make(map[string]*model.Sample),
		order:   []string{},
	}
}

func (s *Store) Record(sm *model.Sample) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.samples[sm.ID]; ok {
		return ErrSampleExists
	}
	s.samples[sm.ID] = sm
	s.order = append(s.order, sm.ID)
	return nil
}

func (s *Store) Get(id string) (*model.Sample, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	sm, ok := s.samples[id]
	if !ok {
		return nil, ErrSampleNotFound
	}
	return sm, nil
}

func (s *Store) ListSamples() []*model.Sample {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]*model.Sample, 0, len(s.order))
	for _, id := range s.order {
		out = append(out, s.samples[id])
	}
	return out
}

func (s *Store) OrderIDs() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]string, len(s.order))
	copy(out, s.order)
	return out
}

func (s *Store) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.order)
}
