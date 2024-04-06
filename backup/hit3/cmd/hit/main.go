package main

import (
	"fmt"
	"os"
	"runtime"
)

const (
	bannerText = `
 __  __     __     ______
/\ \_\ \   /\ \   /\__  _\
\ \  __ \  \ \ \  \/_/\ \/ 
 \ \_\ \_\  \ \_\    \ \_\
  \/_/\/_/   \/_/     \/_/
` // cannot remove first newline character with constant [1:]
)

// Refactor the bannerText and usageText with functions
// This will solve the slicing problem when using const
func banner() string { return bannerText[1:] }

// func usage() string { return usageText[1:] }

func main() {
	// setting a flags object with default values for n and c
	f := &flags{
		n: 100,
		c: runtime.NumCPU(),
	}
	if err := f.parseFlags(); err != nil {
		os.Exit(1)
	}
	fmt.Println(banner())
	fmt.Printf("Making %d requests to %s with a concurrency level of %d.\n", f.n, f.url, f.c)
}
