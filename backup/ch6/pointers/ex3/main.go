package main

// Pointers and interfaces

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct {
}

// method without pointer
func (m martian) talk() string {
	return "nack nack"
}

func main() {
	shout(martian{})
	shout(&martian{})
}
