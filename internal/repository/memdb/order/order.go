package order

import (
	"context"

	"github.com/marshevms/two_gis/internal/repository/model"
)

var actualOrders = []model.Order{}

type Order struct{}

func New() *Order {
	return &Order{}
}

func (o Order) MakeOrder(ctx context.Context, order model.Order) error {
	actualOrders = append(actualOrders, order)

	return nil
}

func (o Order) GetOrders(ctx context.Context, email string) ([]model.Order, error) {
	res := []model.Order{}
	for _, item := range actualOrders {
		if item.Email == email {
			res = append(res, item)
		}
	}

	return res, nil
}
