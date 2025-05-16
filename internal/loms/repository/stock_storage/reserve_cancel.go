package stock_storage

import (
	"context"
	"fmt"
)

func (s *storage) ReserveCancel(ctx context.Context, items []*ReserveItem) error {
	s.Lock()
	defer s.Unlock()

	for _, item := range items {
		storageItem, exists := s.items[item.SkuID]
		if !exists {
			return fmt.Errorf("sku id %d does not exist", item.SkuID)
		}

		if storageItem.Reserved < item.Quantity {
			return fmt.Errorf("reserved %d < %d", storageItem.Reserved, item.Quantity)
		}
	}

	for _, item := range items {
		s.items[item.SkuID].Reserved -= item.Quantity
	}

	return nil
}
