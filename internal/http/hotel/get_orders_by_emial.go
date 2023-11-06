package hotel

import (
	"encoding/json"
	"net/http"
)

func (h *Hotel) GetOrdersByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	orders, err := h.usecase.GetOrdersByEmail(r.Context(), email)
	if err != nil {
		logAndReturnErr("failed to get orders by email", err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(orders); err != nil {
		logAndReturnErr("failed to endcode to json", err, w)
		return
	}

	return
}
