package hotel

import (
	"net/http"
	"time"

	usecase_model "github.com/marshevms/two_gis/internal/usecase/model"
)

func (h *Hotel) MakeOrder(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	room := r.URL.Query().Get("room")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	fromTime, err := time.Parse(time.RFC3339, from)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	toTime, err := time.Parse(time.RFC3339, to)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = h.usecase.MakeOrder(r.Context(), usecase_model.Order{
		Email: email,
		Room:  room,
		From:  fromTime,
		To:    toTime,
	})

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
