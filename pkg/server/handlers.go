package server

import (
	"encoding/json"
	"net/http"

	"github.com/jrmanes/seget/internal/models"
	"github.com/jrmanes/seget/pkg/db/mysql"

	log "github.com/sirupsen/logrus"
)

type Respose struct {
	Item   interface{} `json:"item"` // return any kind of response
	Errors string      `json:"errors"`
}

func GetItemHandler(w http.ResponseWriter, r *http.Request) {

	item := models.Item{}

	item, err := mysql.GetItem()
	if err != nil {
		log.Error("Error marshaling to JSON:", err)
		resp := Respose{
			Errors: string(err.Error()),
			Item:   item,
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

	// Generate the response, adding the size of the array
	resp := Respose{
		Errors: "",
		Item:   item,
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		log.Error("Error marshaling to JSON:", err)
		return
	}

	log.Info(r.Host, " ", r.URL, " ", r.Method, " ", resp)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	item := models.Item{}

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Error("There was an error decoding the request body into the struct")
	}

	err = mysql.AddItem(item)
	if err != nil {
		log.Error("There was an error decoding the request body into the struct")
	}

	// Generate the response, adding the size of the array
	resp := Respose{
		Errors: "",
		Item:   item,
	}

	jsonData, err := json.Marshal(item)
	if err != nil {
		log.Error("Error marshaling to JSON:", err)
		return
	}

	log.Info(r.Host, " ", r.URL, " ", r.Method, " ", resp)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
