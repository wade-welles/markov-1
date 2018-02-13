package storage

import (
	"testing"
)

func TestSlide(t *testing.T) {
	link := Link{Prefix: "I can't", Suffix: "believe"}
	link.Slide()

	if link.Prefix != "can't believe" {
		t.Error("Expected prefix slide to result in `I can't` => `can't believe`.")
	}
}

func TestLoad(t *testing.T) {
	store, err := Load("sqlite3", "./markov.testdb")
	if err != nil {
		t.Error(err.Error())
	}

	if store == nil {
		t.Error("Nonexistent store created.")
	}
}

func TestRegisterDriver(t *testing.T) {
	RegisterDriver("test_sqlite", &SqliteDriver{})

	if drivers["test_sqlite"] == nil {
		t.Error("Test SQLite driver was not added to drivers stack.")
	}
}
