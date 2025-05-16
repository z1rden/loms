package order_storage

import (
	"context"
	"sync"
)

type Storage interface {
	Create(ctx context.Context, userID int64, items []*Item) (int64, error)
	SetStatus(ctx context.Context, userID int64, status string) error
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
