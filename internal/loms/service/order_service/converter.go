package order_service

import (
	"context"
	"loms/internal/loms/model"
	"loms/internal/loms/repository/order_storage"
	"loms/internal/loms/repository/stock_storage"
)

func ToOrderStorageItems(items []*model.Item) []*order_storage.Item {
	res := make([]*order_storage.Item, 0, len(items))
	for _, item := range items {
		res = append(res, toOrderStorageItem(item))
	}

	return res
}

func toOrderStorageItem(item *model.Item) *order_storage.Item {
	return &order_storage.Item{
		SkuID:    item.SkuID,
		Quantity: item.Quantity,
	}
}

func toStockStorageItem(ctx context.Context, item *model.Item) *stock_storage.ReserveItem {
	return &stock_storage.ReserveItem{
		SkuID:    item.SkuID,
		Quantity: item.Quantity,
	}
}

func ToStockStorageItems(ctx context.Context, items []*model.Item) []*stock_storage.ReserveItem {
	reserveItems := make([]*stock_storage.ReserveItem, 0, len(items))
	for _, item := range items {
		reserveItems = append(reserveItems, toStockStorageItem(ctx, item))
	}

	return reserveItems
}
