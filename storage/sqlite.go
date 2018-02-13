package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"sync"
)

type SqliteDriver struct {
	DB    *sql.DB
	Mutex sync.RWMutex
}

func init() {
	RegisterDriver("sqlite3", &SqliteDriver{})
}

func (sd *SqliteDriver) Open(connectionString string) error {
	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		return err
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS markov (prefix TEXT, suffix TEXT)"); err != nil {
		return err
	}

	sd.DB = db
	return nil
}

func (sd *SqliteDriver) AddLink(link Link) error {
	sd.Mutex.Lock()
	defer sd.Mutex.Unlock()

	if _, err := sd.DB.Exec("INSERT INTO markov (prefix, suffix) VALUES ($1, $2)", link.Prefix, link.Suffix); err != nil {
		return err
	}

	return nil
}

func (sd *SqliteDriver) GenerateLine(maxLineLength int) (string, error) {
	var link Link
	var line []string

	query, err := sd.DB.Prepare("SELECT prefix, suffix FROM markov ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		return "", err
	}

	query.QueryRow().Scan(&link.Prefix, &link.Suffix)
	line = append(line, link.Prefix)
	line = append(line, link.Suffix)

	for i := 0; i < maxLineLength-3; i++ {
		link.Slide()
		query, err := sd.DB.Prepare("SELECT suffix FROM markov WHERE prefix = $1 ORDER BY RANDOM() LIMIT 1")
		if err != nil {
			break
		}

		err = query.QueryRow(link.Prefix).Scan(&link.Suffix)
		if err != nil {
			break
		}

		line = append(line, link.Suffix)
	}

	return strings.Join(line, " "), nil
}
