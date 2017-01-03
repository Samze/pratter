package messages

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type AddHandler struct {
	Store MessageStore
}

func (h *AddHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]
	var msg Message

	if r.Body == nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	h.Store.AddMessage(user, msg)
}
