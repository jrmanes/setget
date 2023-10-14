package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"

	"github.com/jrmanes/seget/internal/models"
	log "github.com/sirupsen/logrus"
)

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
	if count == 0 {
		log.Error("ERROR DB is empty, add some values first: ", err)
		return models.Item{}, err
	}

	// generate a random num between 1 and the limit
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(count) + 1

	// GET
	var item models.Item
	err = db.QueryRow("SELECT * FROM "+table+" WHERE id = ?", randomNumber).Scan(&item.ID, &item.Item)
	if err != nil {
		log.Error("ERROR:", err)
		return models.Item{}, err
	}
	defer db.Close()

	return item, nil
}

func AddItem(item models.Item) error {
	db, err := Conn()
	if err != nil {
		log.Error("ERROR:", err)
		return err
	}

	// Create
	_, err = db.Exec("INSERT INTO "+dbName+" (item) VALUES (?)", item.Item)
	if err != nil {
		log.Error("ERROR:", err)
		return err
	}

	return nil
}
