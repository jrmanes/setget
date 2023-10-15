package mysql

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	user   string
	pass   string
	host   string
	port   string
	dbName string
	table  string
)

func Config() {
	user = os.Getenv("MYSQL_USER")
	if user == "" {
		user = ""
	}
	pass = os.Getenv("MYSQL_USER")
	if pass == "" {
		pass = ""
	}
	host = os.Getenv("MYSQL_HOST")
	if host == "" {
		host = "localhost"
	}
	port = os.Getenv("MYSQL_PORT")
	if port == "" {
		port = "3306"
	}
	dbName = os.Getenv("MYSQL_DATABASE")
	if dbName == "" {
		dbName = ""
	}
	table = os.Getenv("MYSQL_TABLE")
	if table == "" {
		table = ""
	}
}

func Conn() (*sql.DB, error) {
	Config()

	connString := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbName

	//db, err := sql.Open("mysql", "user:password@/dbname")
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Error("ERRROR: ", err)
		return db, err
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		log.Error("ERRROR: ", err)
		return db, err
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

func SetupDB() error {
	Config()

	connString := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbName

	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Error("ERRROR: ", err)
		return err
	}

	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		log.Error("ERRROR: ", err)
		panic(err)
	}

	q := "CREATE DATABASE IF NOT EXISTS " + dbName
	log.Info("q: ", q)
	_, err = db.Exec(q)
	if err != nil {
		log.Error("ERRROR: ", err)
		return err
	}

	q2 := "USE " + dbName
	log.Info("q2: ", q2)
	_, err = db.Exec(q2)
	if err != nil {
		log.Error("ERRROR: ", err)
		return err
	}

	query := "CREATE TABLE IF NOT EXISTS " + dbName + " (id int primary key auto_increment, item text)"
	log.Info("query: ", query)

	_, err = db.Exec(query)
	if err != nil {
		log.Error("ERRROR: ", err)
		return err
	}

	return nil
}
