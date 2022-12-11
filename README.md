# Dictionary-CLI

Simple dictionnary CLI using a [badger](https://github.com/dgraph-io/badger) key value database written in go

## Using the CLI

First build the project:

`go build -o build`

Once you've built the project you can start using the CLI as following:

`./build -action add Golang "Golang is a beautiful programming language" `

The previous command allows you to add the word Golang in the dictionnary.

## Commands

Every command is prefixed by the prefix `-action`.
Here a the differents actions available :

- help: `-action help`
- list: `-action list`
- add: `-action add the_word_to_add "the_definition"`
- define: `-action define the_word_to_define`
- remove: `-action remove the_word_to_remove`
