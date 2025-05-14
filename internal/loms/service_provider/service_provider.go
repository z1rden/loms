package service_provider

type ServiceProvider struct {
	repository repository
	service    service
}

var serviceProvider *ServiceProvider

func GetServiceProvider() *ServiceProvider {
	if serviceProvider == nil {
		serviceProvider = &ServiceProvider{}
	}

	return serviceProvider
}
