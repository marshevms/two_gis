package hotel

import (
	"context"
	"fmt"
	"net/mail"

	rep_model "github.com/marshevms/two_gis/internal/repository/model"
	usecase_error "github.com/marshevms/two_gis/internal/usecase/errors"
	"github.com/marshevms/two_gis/internal/usecase/model"
)

type Hotel struct {
	room  Room
	order Order
}

func New(room Room, order Order) *Hotel {
	return &Hotel{
		room:  room,
		order: order,
	}
}

func (h Hotel) MakeOrder(ctx context.Context, order *model.Order) error {
	addr, err := mail.ParseAddress(order.Email)
	if err != nil {
		return fmt.Errorf("failed to parse email '%s': %w: %w", order.Email, usecase_error.InvalidEmail, err)
	}

	order.Email = addr.Address

	if !order.From.Before(order.To) {
		return fmt.Errorf("failed to make order from '%s' to '%s': %w", order.From, order.To, usecase_error.InvalidTimePeriod)
	}

	rooms, err := h.room.GetAvailable(ctx)
	if err != nil {
		return fmt.Errorf("failed to get available rooms: %w", toUsescaseError(err))
	}

	if _, ok := rooms[order.Room]; !ok {
		return fmt.Errorf("failed to make order by room '%s': %w", order.Room, usecase_error.DontHaveAvailableRooms)
	}

	exist, err := h.order.ExistByTime(ctx, order.Email, order.From, order.To)
	if err != nil {
		return toUsescaseError(err)
	}

	if exist {
		return fmt.Errorf("failed to make order from '%s' to '%s': %w", order.From, order.To, usecase_error.OrderForThatTimeAlreadyExist)
	}

	err = h.order.Create(ctx, (*rep_model.Order)(order))
	if err != nil {
		return fmt.Errorf("failed to create order: %w", toUsescaseError(err))
	}

	return nil
}

func (h Hotel) GetOrdersByEmail(ctx context.Context, email string) ([]model.Order, error) {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return nil, fmt.Errorf("failed to parse email '%s': %w: %w", email, usecase_error.InvalidEmail, err)
	}

	email = addr.Address

	orders, err := h.order.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders by email '%s': %w", email, toUsescaseError(err))
	}

	res := make([]model.Order, 0, len(orders))
	for _, order := range orders {
		res = append(res, model.Order(order))
	}

	return res, nil
}
