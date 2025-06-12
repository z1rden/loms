package kafka_storage

import (
	"context"
	"loms/internal/loms/db"
)

type Storage interface {
}

type storage struct {
	ctx      context.Context
	dbClient db.Client
}

func NewStorage(ctx context.Context, dbClient db.Client) Storage {
	return &storage{
		ctx:      ctx,
		dbClient: dbClient,
	}
}
