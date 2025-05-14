package order_service

import (
	"context"
	"loms/internal/loms/repository/order_storage"
)

type Service interface {
	Create()
}

type service struct {
	storage order_storage.Storage
}

func NewService(ctx context.Context, storage order_storage.Storage) Service {
	return service{
		storage: storage,
	}
}
