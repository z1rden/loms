package service_provider

import (
	"context"
	"loms/internal/loms/service/order_service"
)

type service struct {
	orderService order_service.Service
}

func (s *ServiceProvider) GetOrderService(ctx context.Context) order_service.Service {
	if s.service.orderService == nil {
		s.service.orderService = order_service.NewService(ctx, s.GetOrderStorage(ctx))
	}

	return s.service.orderService
}
