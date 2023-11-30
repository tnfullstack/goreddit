// Basic phonebook app
package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name    string
	Surname string
	Phone   string
}

var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	args := os.Args

	if len(args) == 1 {
		exe := path.Base(args[0])
		fmt.Printf("Usage: %s search | list <args>\n", exe)
		return
	}

	data = append(data, Entry{"Mihalis", "Tsoukalos", "2109413456"})
	data = append(data, Entry{"Mary", "Smith", "345213456"})
	data = append(data, Entry{"John", "Doe", "555-665-6968"})

	// Differentiate between the commands
	switch args[1] {
	// The search command
	case "search":
		if len(args) != 3 {
			fmt.Println("Usage search Surname")
			return
		}
		result := search(args[2])
		if result == nil {
			fmt.Println("Entry not found", args[2])
			return
		}
		fmt.Println(*result)
		// The list command
	case "list":
		list()
		// Response to anything that is not a match
	default:
		fmt.Println("Not a valid option")
	}
}
