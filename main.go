package main

import (
	dictionnary "dictionnary/logic"
	"flag"
	"fmt"
	"os"
)

func main() {

	action := flag.String("action", "list", "Action to perform on the dictionnary")

	d, err := dictionnary.New("./badger")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer d.Close()

	flag.Parse()

	switch *action {
	case "help":
		dictionnary.ActionHelp()
	case "list":
		dictionnary.ActionFindAll(d)
	case "add":
		dictionnary.ActionAdd(d, flag.Args())
	case "define":
		dictionnary.ActionDefine(d, flag.Args())
	case "remove":
		dictionnary.ActionRemove(d, flag.Args())
	default:
		fmt.Printf("Unknow action %s please flag your action !\n", *action)
	}

}
