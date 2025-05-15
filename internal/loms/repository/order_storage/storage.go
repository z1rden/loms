package order_storage

import (
	"context"
	"sync"
)

type Storage interface {
	Create(ctx context.Context, userID int64, items []*Item) (int64, error)
}

type storage struct {
	sync.RWMutex
	orders map[int64]*Order
}

func NewStorage(ctx context.Context) Storage {
	return &storage{
		orders: map[int64]*Order{},
	}
}

func (s *storage) getNextID() int64 {
	var maxID int64
	for orderID := range s.orders {
		if orderID > maxID {
			maxID = orderID
		}
	}
	maxID++
	return maxID
}
