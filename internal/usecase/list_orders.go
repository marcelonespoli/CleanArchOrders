package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type ListOrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() ([]ListOrdersOutputDTO, error) {
	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var output []ListOrdersOutputDTO
	for _, o := range orders {
		output = append(output, ListOrdersOutputDTO{
			ID:         o.ID,
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.FinalPrice,
		})
	}

	return output, nil
}
