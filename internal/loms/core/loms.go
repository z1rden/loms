package core

import (
	"context"
	order_api "loms/internal/loms/api/order"
	"loms/internal/loms/grpc_server"
	"loms/internal/loms/http_server"
	"loms/internal/loms/service_provider"
	"loms/pkg/config"
	"loms/pkg/logger"
)

type Service interface {
	Run() error
}

type service struct {
	ctx             context.Context
	serviceProvider *service_provider.ServiceProvider
	cfg             *config.Config
}

func NewService(ctx context.Context) Service {
	ctx, cancel := context.WithCancel(context.Background())
	serviceProvider := service_provider.GetServiceProvider(ctx)
	serviceProvider.GetCloser(ctx).Add(func() error {
		cancel()
		return nil
	})

	cfg := config.MustLoad()
	logger.WithNameApp(ctx, cfg.AppName)

	return &service{
		ctx:             ctx,
		serviceProvider: serviceProvider,
		cfg:             cfg,
	}
}

func (s *service) Run() error {
	logger.Infof(s.ctx, "Starting service")
	defer logger.Infof(s.ctx, "Stopping service")

	closer := s.serviceProvider.GetCloser(s.ctx)
	defer closer.Wait()

	orderApi := s.serviceProvider.GetOrderAPI(s.ctx)

	grpcServer := grpc_server.NewServer(s.ctx, s.cfg.GrpcPort)
	err := grpcServer.RegisterApi([]order_api.API{orderApi})
	if err != nil {
		return err
	}
	closer.Add(grpcServer.Stop)

	go func() {
		logger.Infof(s.ctx, "Starting grpc server on port %s", s.cfg.GrpcPort)
		err := grpcServer.Run()
		if err != nil {
			logger.Fatalf(s.ctx, "grpc server run failed: %s", err)
			closer.CloseAll()
		}
	}()

	httpServer, err := http_server.NewServer(s.ctx, s.cfg.HttpPort, s.cfg.GrpcPort)
	if err != nil {
		return err
	}
	err = httpServer.RegisterApi([]order_api.API{orderApi})
	if err != nil {
		return err
	}
	closer.Add(httpServer.Stop)

	go func() {
		logger.Infof(s.ctx, "Starting http server on port %s", s.cfg.HttpPort)
		err := httpServer.Run()
		if err != nil {
			logger.Fatalf(s.ctx, "http server run failed: %s", err)
			closer.CloseAll()
		}
	}()

	closer.Add(logger.Close)

	return nil
}
