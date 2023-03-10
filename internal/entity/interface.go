package entity

type OrderRepository interface {
	Save(order *Order) error
	GetTotal() (int, error)
}
