package order_service

import (
	"context"
	"loms/internal/loms/model"
)

func (s *service) Create(ctx context.Context, userID int64, items []*model.Item) (int64, error) {
	orderID, err := s.storage.Create(ctx, userID, ToOrderStorageItems(items))
	if err != nil {
		return 0, err
	}

	return orderID, nil
}
