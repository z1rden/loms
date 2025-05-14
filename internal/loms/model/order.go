package model

const (
	OrederStatusNew            = "new"
	OrederStatusAwatingPayment = "awaiting_payment"
	OrederStatusPayed          = "payed"
	OrederStatusCanceled       = "canceled"
	OrederStatusFailed         = "failed"
)

type Item struct {
	SkuID    int64
	Quantity uint16
}

type Order struct {
	OrderID int64
	User    int64
	Status  string
	Items   []*Item
}

type Orders map[int64]*Order
