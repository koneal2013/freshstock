package store

import (
	"errors"
	"strings"
	"sync"

	"github.com/koneal2013/freshstock/internal/model"
)

var (
	// ErrProduceExists is an error returned when a produce with the given code already exists.
	ErrProduceExists = errors.New("produce with this code already exists")
	// ErrProduceNotFound is an error returned when a produce is not found.
	ErrProduceNotFound = errors.New("produce not found")
)

// ProduceStore is a type that represents a store for produce items.
type ProduceStore struct {
	sync.RWMutex
	produce map[string]*model.Produce
}

// NewProduceStore creates a new instance of ProduceStore.
func NewProduceStore() *ProduceStore {
	return &ProduceStore{
		produce: make(map[string]*model.Produce),
	}
}

// AddProduce adds a new produce to the ProduceStore.
func (s *ProduceStore) AddProduce(produce *model.Produce) error {
	s.Lock()
	defer s.Unlock()

	if _, exists := s.produce[produce.Code]; exists {
		return ErrProduceExists
	}

	s.produce[produce.Code] = produce
	return nil
}

// GetProduceByCode retrieves the produce with the specified code from the ProduceStore.
//
// If the code does not exist in the ProduceStore, it returns nil and an error of type ErrProduceNotFound.
// Otherwise, it returns the produce and a nil error.
func (s *ProduceStore) GetProduceByCode(code string) (*model.Produce, error) {
	s.RLock()
	defer s.RUnlock()

	produce, exists := s.produce[code]
	if !exists {
		return nil, ErrProduceNotFound
	}

	return produce, nil
}

// SearchProduce searches for produce by name or code and returns the matching results.
func (s *ProduceStore) SearchProduce(query string) []*model.Produce {
	s.RLock()
	defer s.RUnlock()

	var results []*model.Produce
	for _, produce := range s.produce {
		if strings.Contains(strings.ToLower(produce.Name), strings.ToLower(query)) {
			results = append(results, produce)
		}
	}

	return results
}

// DeleteProduce deletes a produce item from the store based on its code
func (s *ProduceStore) DeleteProduce(code string) error {
	s.Lock()
	defer s.Unlock()

	if _, exists := s.produce[code]; !exists {
		return ErrProduceNotFound
	}

	delete(s.produce, code)
	return nil
}
