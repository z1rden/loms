package order_storage

import (
	"context"
	"loms/internal/loms/model"
)

func (s *storage) Create(ctx context.Context, userID int64, items []*Item) (int64, error) {
	s.Lock()
	defer s.Unlock()

	order := &Order{
		OrderID: s.getNextID(),
		UserID:  userID,
		Items:   items,
		Status:  model.OrderStatusNew,
	}

	s.orders[order.OrderID] = order

	return order.OrderID, nil
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
