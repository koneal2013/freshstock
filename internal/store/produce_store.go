package store

import (
	"sync"

	"github.com/koneal2013/freshstock/internal/model"
)

type ProduceStore struct {
	sync.RWMutex
	produce map[string]*model.Produce
}

func NewProduceStore() *ProduceStore {
	return &ProduceStore{
		produce: make(map[string]*model.Produce),
	}
}

func (s *ProduceStore) AddProduce(m *model.Produce) error {
	return nil
}

func (s *ProduceStore) GetProduceByCode(code string) (*model.Produce, error) {
	return nil, nil
}

func (s *ProduceStore) SearchProduce(query string) []*model.Produce {
	return nil
}

func (s *ProduceStore) DeleteProduce(code string) error {
	return nil
}
