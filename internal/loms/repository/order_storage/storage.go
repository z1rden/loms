package order_storage

import (
	"context"
	"sync"
)

type Storage interface {
	Create(ctx context.Context, orderID int64, items []*Item) (int64, error)
	SetStatus(ctx context.Context, orderID int64, status string) error
	GetByID(ctx context.Context, orderID int64) (*Order, error)
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
