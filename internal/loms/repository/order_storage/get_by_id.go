package order_storage

import (
	"context"
	"loms/internal/loms/model"
)

func (s *storage) GetByID(ctx context.Context, orderID int64) (*Order, error) {
	s.RLock()
	defer s.RUnlock()

	order, exists := s.orders[orderID]
	if !exists {
		return nil, model.ErrNotFound
	}

	return order, nil
}
