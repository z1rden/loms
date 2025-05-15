package order_storage

type Order struct {
	OrderID int64
	UserID  int64
	Items   []*Item
	Status  string
}

type Item struct {
	SkuID    int64
	Quantity uint16
}
