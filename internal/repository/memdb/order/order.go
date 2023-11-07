package order

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/marshevms/two_gis/internal/repository/model"
)

var actualOrders = []model.Order{}
var mx = sync.RWMutex{}

var id = atomic.Int64{}

type Order struct{}

func New() *Order {
	return &Order{}
}

func (o Order) Create(ctx context.Context, order *model.Order) error {
	order.ID = id.Add(1)

	mx.Lock()
	defer mx.Unlock()
	actualOrders = append(actualOrders, *order)

	return nil
}

func (o Order) GetByEmail(ctx context.Context, email string) ([]model.Order, error) {
	mx.RLock()
	defer mx.RUnlock()

	res := []model.Order{}
	for _, item := range actualOrders {
		if item.Email == email {
			res = append(res, item)
		}
	}

	return res, nil
}

func (o Order) ExistByTime(ctx context.Context, email string, from, to time.Time) (bool, error) {
	mx.RLock()
	defer mx.RUnlock()

	for _, item := range actualOrders {
		if item.Email != email {
			continue
		}

		if isTimeRangesIntersect([2]time.Time{item.From, item.To}, [2]time.Time{from, to}) {
			return true, nil
		}
	}

	return false, nil
}

func isTimeRangesIntersect(timeRange1 [2]time.Time, timeRange2 [2]time.Time) bool {
	if timeRange1[1].Before(timeRange2[0]) || timeRange1[0].After(timeRange2[1]) {
		return false
	}
	return true
}
