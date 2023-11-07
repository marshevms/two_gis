package hotel

import (
	"context"
	"time"

	"github.com/marshevms/two_gis/internal/repository/model"
)

type Room interface {
	GetAvailable(ctx context.Context) (map[string]struct{}, error)
}

type Order interface {
	Create(ctx context.Context, order *model.Order) error
	GetByEmail(ctx context.Context, email string) ([]model.Order, error)
	ExistByTime(ctx context.Context, email string, from time.Time, to time.Time) (bool, error)
}
