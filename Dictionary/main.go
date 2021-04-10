package main

import (
	"flag"
	"fmt"
	"os"

	"training.go/Dictionary/dictionary"
)

func main() {

	action := flag.String("action", "list", "action to perform on the dictionary")

	d, err := dictionary.New("./badger")
	handleErr(err)
	defer d.Close()

	flag.Parse()
	switch *action {
	case "list":
		actionList(d)
	case "add":
		actionAdd(d, flag.Args())
	case "define":
		actionDefine(d, flag.Arg(0))
	case "remove":
		actionRemove(d, flag.Arg(0))
	default:
		fmt.Printf("Unknown action: %v\n", *action)
	}
}

func actionRemove(d *dictionary.Dictionary, arg string) {
	err := d.Remove(arg)
	handleErr(err)
	fmt.Printf("'%v' was removed from the dictionary\n", arg)
}

func actionDefine(d *dictionary.Dictionary, arg string) {
	entry, err := d.Get(arg)
	handleErr(err)
	fmt.Println(entry)
}

func actionAdd(d *dictionary.Dictionary, args []string) {
	word := args[0]
	definition := args[1]

	err := d.Add(word, definition)
	handleErr(err)
	fmt.Printf("%v added to the dictionary\n", word)
}

func actionList(d *dictionary.Dictionary) {
	words, entries, err := d.List()
	handleErr(err)
	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error:%v\n", err)
		os.Exit(1)
	}
}
