package messages

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ListHandler struct {
	Store *MessageStore
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]

	userMessages := h.Store.getMessages(user)

	if err := json.NewEncoder(w).Encode(&userMessages); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
