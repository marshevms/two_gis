package order

import (
	"context"
	"sync/atomic"

	"github.com/marshevms/two_gis/internal/repository/model"
)

var actualOrders = []model.Order{}
var id = atomic.Int64{}

type Order struct{}

func New() *Order {
	return &Order{}
}

func (o Order) Create(ctx context.Context, order *model.Order) error {
	order.ID = id.Add(1)
	actualOrders = append(actualOrders, *order)

	return nil
}

func (o Order) GetByEmail(ctx context.Context, email string) ([]model.Order, error) {
	res := []model.Order{}
	for _, item := range actualOrders {
		if item.Email == email {
			res = append(res, item)
		}
	}

	return res, nil
}
