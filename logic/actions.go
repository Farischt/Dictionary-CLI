package logic

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"
	"time"

	"github.com/dgraph-io/badger/v3"
)

// Add adds a word to the dictionnary
func (d *Dictionnary) Add(word string, definition string) error {

	entry := Entry{
		Word:       strings.Title(word),
		Definition: definition,
		CreatedAt:  time.Now(),
	}

	var buffer bytes.Buffer
	var encoder *gob.Encoder = gob.NewEncoder(&buffer)
	err := encoder.Encode(entry)

	if err != nil {
		return err
	}

	return d.database.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(word), buffer.Bytes())
	})
}

func (d *Dictionnary) GetDefinition(word string) (Entry, error) {

	var entry Entry
	err := d.database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(strings.Title(word)))

		if err != nil {
			return err
		}

		entry, err = getEntry(item)
		return err
	})

	return entry, err

}

func (d *Dictionnary) FindAll() ([]string, map[string]Entry, error) {

	entries := make(map[string]Entry)

	err := d.database.View(func(txn *badger.Txn) error {
		options := badger.DefaultIteratorOptions
		options.PrefetchSize = 20
		iterator := txn.NewIterator(options)
		defer iterator.Close()

		for iterator.Rewind(); iterator.Valid(); iterator.Next() {
			item := iterator.Item()
			entry, err := getEntry(item)

			if err != nil {
				fmt.Println("Error while getting an entry...")
			}
			entries[entry.Word] = entry

		}

		return nil
	})

	keys := getSortedKeys(entries)

	if err != nil {
		return keys, entries, err
	}
	return keys, entries, nil

}

func (d *Dictionnary) Remove(word string) error {
	var key = []byte(strings.Title(word))
	return d.database.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}
