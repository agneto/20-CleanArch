package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	ListAllOrders() ([]Order, error)
	// GetTotal() (int, error)
}
