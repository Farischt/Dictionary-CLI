package logic

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/dgraph-io/badger/v3"
)

// If error is not nil
// Print the error and exit
func HandleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func ActionFindAll(d *Dictionnary) {
	words, entries, err := d.FindAll()
	HandleError(err)

	fmt.Println("--------- Dictionnary content ---------")
	for _, element := range words {
		fmt.Println(entries[element])
	}
}

func ActionAdd(d *Dictionnary, args []string) {
	if len(args) < 2 {
		HandleError(errors.New("missing word and definition"))
	}

	word := args[0]
	definition := args[1]

	err := d.Add(word, definition)
	HandleError(err)

	fmt.Printf("Word %s added !\n", word)
}

func ActionDefine(d *Dictionnary, args []string) {
	if len(args) < 1 {
		HandleError(errors.New("missing word to define"))
	}
	wordToDefine := args[0]

	entry, err := d.GetDefinition(wordToDefine)
	HandleError(err)

	fmt.Printf("%s: %s\n", entry.Word, entry.Definition)
}

func ActionRemove(d *Dictionnary, args []string) {
	if len(args) < 1 {
		HandleError(errors.New("missing word to remove"))
	}
	wordToRemove := args[0]

	entry, err := d.GetDefinition(wordToRemove)
	HandleError(err)

	err = d.Remove(entry.Word)
	HandleError(err)

	fmt.Printf("Word %s removed !\n", entry.Word)
}

func ActionHelp() {
	fmt.Println("Here are the flags you can use !")
	fmt.Println("-action help")
	fmt.Println(`-action add : with 2 parameters, the word, and its "definition"`)
	fmt.Println("-action remove : with 1 parameter, the word to remove")
	fmt.Println("-action define : with 1 parameter, the word to define")
	fmt.Println("-action list")
}

func getEntry(item *badger.Item) (Entry, error) {
	var entry Entry

	var buffer bytes.Buffer
	err := item.Value(func(val []byte) error {
		_, err := buffer.Write(val)
		return err
	})

	if err != nil {
		return entry, err
	}

	var decoder *gob.Decoder = gob.NewDecoder(&buffer)
	err = decoder.Decode(&entry)

	return entry, err

}

func getSortedKeys(entries map[string]Entry) []string {

	keys := make([]string, len(entries))

	for key := range entries {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}
