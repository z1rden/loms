package stock_storage

type Item struct {
	SkuID      int64  `json:"sku"`
	TotalCount uint16 `json:"total_count"`
	Reserved   uint16 `json:"reserved"`
}

type ReserveItem struct {
	SkuID    int64
	Quantity uint16
}
