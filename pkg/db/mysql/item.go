package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"

	"github.com/jrmanes/seget/internal/models"

	log "github.com/sirupsen/logrus"
)

// GetItem retrieves a random item from the database.
func GetItem() (models.Item, error) {
	db, err := Conn()
	if err != nil {
		log.Error("ERROR:", err)
		return models.Item{}, err
	}

	// Get the number of items in the table
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM " + table).Scan(&count)
	if err != nil {
		log.Error("ERROR:", err)
		return models.Item{}, err
	}

	// verify that we have at least one item in the db
	if count == 0 {
		log.Error("ERROR DB is empty, add some values first: ", err)
		return models.Item{}, err
	}

	// Generate a random number between 1 and the limit
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(count) + 1

	// Retrieve the item with the generated random ID
	var item models.Item
	err = db.QueryRow("SELECT * FROM "+table+" WHERE id = ?", randomNumber).Scan(&item.ID, &item.Item)
	if err != nil {
		log.Error("ERROR:", err)
		return models.Item{}, err
	}
	defer db.Close()

	return item, nil
}

// AddItem adds an item to the database.
func AddItem(item models.Item) error {
	db, err := Conn()
	if err != nil {
		log.Error("ERROR:", err)
		return err
	}

	// Insert the new item into the database
	_, err = db.Exec("INSERT INTO "+dbName+" (item) VALUES (?)", item.Item)
	if err != nil {
		log.Error("ERROR:", err)
		return err
	}

	return nil
}
