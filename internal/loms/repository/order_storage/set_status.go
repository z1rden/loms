package order_storage

import (
	"context"
	"loms/internal/loms/model"
)

func (s *storage) SetStatus(ctx context.Context, orderID int64, status string) error {
	s.Lock()
	defer s.Unlock()

	order, exists := s.orders[orderID]
	if !exists {
		return model.ErrNotFound
	}

	order.Status = status

	return nil

}
