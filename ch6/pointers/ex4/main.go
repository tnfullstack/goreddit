package main

// Pointers and Interfaces

import (
	"fmt"
	"strings"
)

type laser int

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

// method with pointer
func (l *laser) talk() string {
	return strings.Repeat("pew ", int(*l))
}

func main() {
	pew := laser(2)
	fmt.Printf("%v, %T\n", pew, pew)
	shout(&pew)
}
