package service_provider

import (
	"context"
	order_api "loms/internal/loms/api/order"
)

type api struct {
	orderAPI order_api.API
}

func (s *ServiceProvider) GetOrderAPI(ctx context.Context) order_api.API {
	if s.api.orderAPI == nil {
		s.api.orderAPI = order_api.NewApi(ctx, s.GetOrderService(ctx))
	}

	return s.api.orderAPI
}
