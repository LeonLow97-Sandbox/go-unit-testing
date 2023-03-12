package main

import (
	"os"
	"testing"
)

var app application

// Set up Test Environment
func TestMain(m *testing.M) {
	pathToTemplates = "./../../templates/"

	// Setting up test environment for sessions
	app.Session = getSession()

	os.Exit(m.Run())
}
