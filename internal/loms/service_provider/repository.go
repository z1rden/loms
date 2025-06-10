package service_provider

import (
	"context"
	"loms/internal/loms/db"
	"loms/internal/loms/logger"
	"loms/internal/loms/repository/order_storage"
	"loms/internal/loms/repository/stock_storage"
)

type repository struct {
	dbCLient     db.Client
	orderStorage order_storage.Storage
	stockStorage stock_storage.Storage
}

func (s *ServiceProvider) GetDBClient(ctx context.Context) db.Client {
	if s.repository.dbCLient == nil {
		var err error
		s.repository.dbCLient, err = db.NewClient(ctx, s.cfg.MasterDBURL, s.cfg.SyncDBURL)
		if err != nil {
			logger.Fatalf(ctx, "failed to create db client: %v", err)
		}

		s.GetCloser(ctx).Add(s.repository.dbCLient.Close)
	}

	return s.repository.dbCLient
}

func (s *ServiceProvider) GetOrderStorage(ctx context.Context) order_storage.Storage {
	if s.repository.orderStorage == nil {
		s.repository.orderStorage = order_storage.NewStorage(ctx)
	}

	return s.repository.orderStorage
}

func (s *ServiceProvider) GetStockStorage(ctx context.Context) stock_storage.Storage {
	if s.repository.stockStorage == nil {
		s.repository.stockStorage = stock_storage.NewStorage(ctx, s.GetDBClient(ctx))
	}

	return s.repository.stockStorage
}
