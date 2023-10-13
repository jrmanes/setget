package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Respose struct {
	Size   int `json:"size"`
	Errors int `json:"errors"`
}

func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	// get the user by param
	user_id := mux.Vars(r)["user"]
	if user_id == "" {
		log.Error("User param is empty", http.StatusUnauthorized)
		return
	}

	// Generate the response, adding the size of the array
	resp := Respose{
		Errors: 0,
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		log.Error("Error marshaling to JSON:", err)
		return
	}

	//w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	// err = json.NewDecoder(r.Body).Decode(&location)
	// if err != nil {
	// 	log.Error("There was an error decoding the request body into the struct")
	// }

	// uuid := uuid.New().String()
	// location.Id = uuid

	jsonData, err := json.Marshal(location)
	if err != nil {
		log.Error("Error marshaling to JSON:", err)
		return
	}

	//w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
