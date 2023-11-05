package hotel

import (
	"encoding/json"
	"net/http"
)

func (h *Hotel) GetOrdersByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	orders, err := h.usecase.GetOrdersByEmail(r.Context(), email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(orders); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
