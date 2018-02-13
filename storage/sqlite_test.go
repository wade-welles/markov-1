package storage

import (
	"testing"
)

var sd SqliteDriver

func TestDial(t *testing.T) {
	if err := sd.Open("./markov_test.db"); err != nil {
		t.Error(err.Error())
	}

	if err := sd.DB.Ping(); err != nil {
		t.Error(err.Error())
	}
}