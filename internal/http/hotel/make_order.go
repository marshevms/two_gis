package hotel

import (
	"fmt"
	"net/http"
	"time"

	usecase_model "github.com/marshevms/two_gis/internal/usecase/model"
)

func (h Hotel) MakeOrder(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	room := r.URL.Query().Get("room")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	fromTime, err := time.Parse(time.RFC3339, from)
	if err != nil {
		logAndReturnCode(fmt.Sprintf("failed to parse from: '%s'", from), http.StatusBadRequest, w)
		return
	}
	toTime, err := time.Parse(time.RFC3339, to)
	if err != nil {
		logAndReturnCode(fmt.Sprintf("failed to parse to: '%s'", to), http.StatusBadRequest, w)
		return
	}

	order := &usecase_model.Order{
		Email: email,
		Room:  room,
		From:  fromTime,
		To:    toTime,
	}

	err = h.usecase.MakeOrder(r.Context(), order)
	if err != nil {
		logAndReturnErr("failed to make order", err, w)
		return
	}
}
