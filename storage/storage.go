package storage

import (
	"errors"
	"fmt"
)

var drivers = make(map[string]Storage)

type Link struct {
	Prefix string
	Suffix string
}

type Storage interface {
	Open(string) error
}

type Connection interface {
	Close() error
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
