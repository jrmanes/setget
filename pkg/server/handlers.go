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

// Response represents the JSON response structure.
type Response struct {
	Item   interface{} `json:"item"` // Return any type
	Errors string      `json:"errors"`
}

// AddItemHandler handles HTTP POST requests for adding an item.
func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	item := models.Item{}

	// Decode the request body into the `item` struct.
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Error("There was an error decoding the request body into the struct.")
	}

	// check if the item received is empty
	if item.Item == "" {
		ResponseHttp(w, r, item, errors.New("ERROR: item is empty"))
		log.Error("ERROR: item is empty.")
	}

	// Call the `mysql.AddItem` function to add an item.
	err = mysql.AddItem(item)
	if err != nil {
		log.Error("There was an error decoding the request body into the struct.")
	}

	// Generate the response, adding the size of the array.
	ResponseHttp(w, r, item, err)
}

// GetItemHandler handles HTTP GET requests for retrieving an item.
func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	item := models.Item{}

	// Call the `mysql.GetItem` function to retrieve an item.
	item, err := mysql.GetItem()
	if err != nil {
		log.Error("Error marshaling to JSON:", err)
		ResponseHttp(w, r, item, err)
	}

	ResponseHttp(w, r, item, err)
}

// ResponseHttp generates an HTTP response with JSON content.
func ResponseHttp(w http.ResponseWriter, r *http.Request, item models.Item, err error) {
	resp := Response{}

	if err != nil {
		resp = Response{
			Errors: fmt.Sprintf("ERROR: %d", err),
			Item:   item,
		}
	} else {
		resp = Response{
			Errors: "",
			Item:   item,
		}
	}

	// Marshal the `item` into JSON.
	jsonData, err := json.Marshal(item)

	if err != nil {
		log.Error("Error marshaling to JSON:", err)
		return
	}

	// Log information about the request.
	log.Info(r.Host, " ", r.URL, " ", r.Method, " ", resp)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
