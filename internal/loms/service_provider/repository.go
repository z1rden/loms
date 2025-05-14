package service_provider

import (
	"context"
	"loms/internal/loms/repository/order_storage"
)

type repository struct {
	orderStorage order_storage.Storage
}

func (s *ServiceProvider) GetOrderStorage(ctx context.Context) order_storage.Storage {
	if s.repository.orderStorage == nil {
		s.repository.orderStorage = order_storage.NewStorage(ctx)
	}

	return s.repository.orderStorage
}
