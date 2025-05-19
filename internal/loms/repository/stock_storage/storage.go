package stock_storage

import (
	"context"
	"encoding/json"
	"loms/internal/loms/logger"
	"os"
	"sync"
)

type Storage interface {
	Reserve(ctx context.Context, items []*ReserveItem) error
	GetBySku(ctx context.Context, SkuID int64) (uint16, error)
	ReserveCancel(ctx context.Context, items []*ReserveItem) error
	ReserveRemove(ctx context.Context, items []*ReserveItem) error
}

type storage struct {
	sync.RWMutex
	items map[int64]*Item
}

func NewStorage(ctx context.Context) Storage {
	s := &storage{
		items: map[int64]*Item{},
	}

	jsonFilePath := "./config/stock_data.json"
	if _, err := os.Stat(jsonFilePath); os.IsNotExist(err) {
		logger.Errorf(ctx, "json file path does not exist: %s", jsonFilePath)
	} else {
		fileData, err := os.ReadFile(jsonFilePath)
		if err != nil {
			logger.Errorf(ctx, "Read json file error: %s", err)
		}

		var items []*Item
		if err := json.Unmarshal(fileData, &items); err != nil {
			logger.Errorf(ctx, "Unmarshal json file error: %s", err)
		}

		for _, item := range items {
			s.items[item.SkuID] = item
		}
	}

	return s
}
