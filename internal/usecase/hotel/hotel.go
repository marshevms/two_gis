package hotel

import (
	"context"
	"fmt"

	rep_model "github.com/marshevms/two_gis/internal/repository/model"
	usecase_error "github.com/marshevms/two_gis/internal/usecase/errors"
	"github.com/marshevms/two_gis/internal/usecase/model"
)

type Hotel struct {
	room  Room
	order Order
}

func New() *Hotel {
	return &Hotel{}
}

func (h *Hotel) MakeOrder(ctx context.Context, order model.Order) error {
	rooms, err := h.room.GetAvailable(ctx)
	if err != nil {
		return err
	}

	if _, ok := rooms[order.Room]; !ok {
		return fmt.Errorf("failed to make order by room %s: %w", order.Room, usecase_error.DontHaveAvailableRooms)
	}

	exist, err := h.order.ExistByTime(ctx, order.From, order.To)
	if err != nil {
		return err
	}

	if !exist {
		return fmt.Errorf("failed to make order from %s to %s: %w", order.From, order.To, usecase_error.OrderForThatTimeAlreadyExist)
	}

	return h.order.Make(ctx, rep_model.Order(order))
}

func (h *Hotel) GetOrdersByEmail(ctx context.Context, userEmail string) ([]model.Order, error) {
	orders, err := h.order.GetByEmail(ctx, userEmail)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders by email %s: %w", userEmail, err)
	}

	res := make([]model.Order, 0, len(orders))
	for _, order := range orders {
		res = append(res, model.Order(order))
	}

	return res, err
}
