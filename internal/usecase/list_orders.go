package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type ListOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	ListOrder       events.EventInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	ListOrder events.EventInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
		ListOrder:       ListOrder,
	}
}

func (c *ListOrderUseCase) Execute() ([]ListOrderOutputDTO, error) {
	orders, err := c.OrderRepository.ListAllOrders()
	if err != nil {
		return []ListOrderOutputDTO{}, err
	}

	var list []ListOrderOutputDTO
	for _, order := range orders {
		var dto = ListOrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		}
		list = append(list, dto)
	}

	return list, nil
}
