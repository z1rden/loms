package order_storage

import "context"

func (s *storage) Create(ctx context.Context, userID int64, items []*Item) (int64, error) {
	s.Lock()
	defer s.Unlock()

	order := &Order{
		OrderID: s.getNextID(),
		UserID:  userID,
		Items:   items,
		Status:  "new",
	}

	s.orders[order.OrderID] = order

	return order.OrderID, nil
}
