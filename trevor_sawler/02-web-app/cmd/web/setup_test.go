package main

import (
	"log"
	"os"
	"testing"
	"webapp/pkg/db"
)

var app application

// Set up Test Environment
func TestMain(m *testing.M) {
	pathToTemplates = "./../../templates/"

	// Setting up test environment for sessions
	app.Session = getSession()

	app.DSN = "host=localhost port=5433 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5"

	// For database connection
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.DB = db.PostgresConn{DB: conn}

	os.Exit(m.Run())
}
