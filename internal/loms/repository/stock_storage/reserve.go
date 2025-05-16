package stock_storage

import (
	"context"
	"errors"
	"loms/internal/loms/model"
)

func (s *storage) Reserve(ctx context.Context, items []*ReserveItem) error {
	s.Lock()
	defer s.Unlock()

	for _, item := range items {
		storageItem, ok := s.items[item.SkuID]
		if !ok {
			return model.ErrNotFound
		}

		if storageItem.TotalCount-storageItem.Reserved < item.Quantity {
			return errors.New("reservation limit reached")
		}
	}

	for _, item := range items {
		s.items[item.SkuID].Reserved += item.Quantity
	}

	return nil
}
