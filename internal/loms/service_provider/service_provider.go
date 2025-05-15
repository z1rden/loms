package service_provider

import (
	"context"
	"loms/pkg/closer"
)

type ServiceProvider struct {
	repository repository
	service    service
	closer     closer.Closer
	api        api
}

var serviceProvider *ServiceProvider

func GetServiceProvider(ctx context.Context) *ServiceProvider {
	if serviceProvider == nil {
		serviceProvider = &ServiceProvider{}
	}

	return serviceProvider
}
