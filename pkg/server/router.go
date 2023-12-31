package server

import (
	"github.com/gorilla/mux"
)

func Router(r *mux.Router) *mux.Router {
	r.HandleFunc("/set", AddItemHandler).Methods("POST")
	r.HandleFunc("/get", GetItemHandler).Methods("GET")

	return r
}
