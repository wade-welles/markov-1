package storage

import (
	"testing"
)

var sd SqliteDriver

func TestOpen(t *testing.T) {
	if err := sd.Open("./markov.testdb"); err != nil {
		t.Error(err.Error())
	}

	if err := sd.DB.Ping(); err != nil {
		t.Error(err.Error())
	}
}

func TestAddLink(t *testing.T) {
	link := Link{Prefix: "I can't", Suffix: "believe"}

	if err := sd.AddLink(link); err != nil {
		t.Error(err.Error())
	}
}

func TestGenerateLine(t *testing.T) {
	if _, err := sd.GenerateLine(6); err != nil {
		t.Error(err.Error())
	}
}
