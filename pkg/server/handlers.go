package server

import (
	"encoding/json"
	"errors"
	"fmt"
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
		ResponseHttp(w, r, item, err)
	}

	// Generate the response, adding the size of the array
	resp := Respose{
		Errors: "",
		Item:   item,
	}

	log.Info(r.Host, " ", r.URL, " ", r.Method, " ", resp)

	ResponseHttp(w, r, item, err)
}

func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	item := models.Item{}

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Error("There was an error decoding the request body into the struct")
	}

	if item.Item == "" {
		ResponseHttp(w, r, item, errors.New("ERROR: item is emtpy"))
		log.Error("ERROR: item is emtpy")
	}

	err = mysql.AddItem(item)
	if err != nil {
		log.Error("There was an error decoding the request body into the struct")
	}

	// Generate the response, adding the size of the array
	ResponseHttp(w, r, item, err)
}

func ResponseHttp(w http.ResponseWriter, r *http.Request, item models.Item, err error) {
	resp := Respose{}
	if err != nil {
		resp = Respose{
			Errors: fmt.Sprintf("ERROR: %d", err),
			Item:   item,
		}
	} else {
		resp = Respose{
			Errors: "",
			Item:   item,
		}
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
