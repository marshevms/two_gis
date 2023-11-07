package hotel

import (
	"context"

	usecase_model "github.com/marshevms/two_gis/internal/usecase/model"
)

type Usecase interface {
	GetOrdersByEmail(ctx context.Context, userEmail string) ([]usecase_model.Order, error)
	MakeOrder(ctx context.Context, order *usecase_model.Order) error
}

type Hotel struct {
	usecase Usecase
}

func New(usecase Usecase) *Hotel {
	return &Hotel{
		usecase: usecase,
	}
}
