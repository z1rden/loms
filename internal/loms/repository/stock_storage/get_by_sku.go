package stock_storage

import (
	"context"
	"loms/internal/loms/model"
)

func (s *storage) GetBySku(ctx context.Context, SkuID int64) (uint16, error) {
	s.RLock()
	defer s.RUnlock()

	item, ok := s.items[SkuID]
	if !ok {
		return 0, model.ErrNotFound
	}

	return item.TotalCount - item.Reserved, nil

}
