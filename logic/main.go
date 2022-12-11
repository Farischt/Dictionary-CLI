package logic

import (
	"fmt"
	"time"

	"github.com/dgraph-io/badger/v3"
)

const IsModeDebug bool = false

type Dictionnary struct {
	database *badger.DB
}

type Entry struct {
	Word       string
	Definition string
	CreatedAt  time.Time
}

func (e Entry) String() string {
	created := e.CreatedAt.Format(time.Stamp)
	return fmt.Sprintf("%-10s \t %-50s \t %-6v", e.Word, e.Definition, created)
}

// New creates a new dictionnary with a database
// Takes a string (the dir path to store db) as parameter
// Returns a pointer to the Dictonnary and an error
func New(dir string) (*Dictionnary, error) {
	options := badger.DefaultOptions(dir)

	if !IsModeDebug {
		options.Logger = nil
	}

	db, err := badger.Open(options)
	if err != nil {
		return nil, err
	}

	dict := &Dictionnary{
		database: db,
	}
	return dict, nil
}

// Close closes the dictionnary database
func (d *Dictionnary) Close() {
	d.database.Close()
}
