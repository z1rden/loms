package order_service

import (
	"loms/internal/loms/model"
	"loms/internal/loms/repository/order_storage"
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
