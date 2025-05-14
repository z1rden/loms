package core

import (
	"context"
	"errors"
	"loms/internal/loms/service_provider"
)

type Service interface {
	Run() error
}

type service struct {
	ctx             context.Context
	serviceProvider service_provider.ServiceProvider
}

func NewService(ctx context.Context) Service {
	return &service{
		ctx: ctx,
	}
}

func (s *service) Run() error {
	return errors.New("not implemented")
}
