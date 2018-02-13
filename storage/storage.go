package storage

import (
	"errors"
	"fmt"
	"strings"
)

var drivers = make(map[string]Storage)

type Link struct {
	Prefix string
	Suffix string
}

type Storage interface {
	Open(string) error
	GenerateLine(int) (string, error)
	AddLink(Link) error
}

func (link *Link) Slide() {
	prefixWords := strings.Split(link.Prefix, " ")
	var newPrefix []string
	newPrefix = append(newPrefix, prefixWords[1])
	newPrefix = append(newPrefix, link.Suffix)
	link.Prefix = strings.Join(newPrefix, " ")
}

func Load(driverName string, connectionString string) (Storage, error) {
	if drivers[driverName] == nil {
		err := errors.New(fmt.Sprintf("Could not load a storage driver with name %s.", driverName))
		return nil, err
	}

	driver := drivers[driverName]
	if err := driver.Open(connectionString); err != nil {
		return nil, err
	}

	return driver, nil
}

func RegisterDriver(driverName string, driver Storage) error {
	drivers[driverName] = driver
	return nil
}
