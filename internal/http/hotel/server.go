package hotel

import (
	"context"

	usecase_model "github.com/marshevms/two_gis/internal/usecase/model"
)

type Hotel struct {
	usecase Usecase
}

type Usecase interface {
	GetOrdersByEmail(ctx context.Context, userEmail string) ([]usecase_model.Order, error)
	MakeOrder(ctx context.Context, order usecase_model.Order) error
}
