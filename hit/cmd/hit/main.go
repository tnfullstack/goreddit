package main

import (
	"fmt"
	"log"
)

const (
	bannerText = `
 __  __     __     ______
/\ \_\ \   /\ \   /\__  _\
\ \  __ \  \ \ \  \/_/\ \/ 
 \ \_\ \_\  \ \_\    \ \_\
  \/_/\/_/   \/_/     \/_/
` // cannot remove first newline character with constant [1:]

	usageText = `
Usage:
	-url   - HTTP server URL to make request (required)
	-n     - Number of requests to make
	-c     - Concurrency level` // cannot use slice [1:] in constant
)

// Refactor the bannerText and usageText with functions
// This will solve the slicing problem when using const
func banner() string { return bannerText[1:] }
func usage() string  { return usageText[1:] }

func main() {
	var f flags
	if err := f.parse(); err != nil {
		fmt.Println(usage())
		log.Fatal(err)
	}
	fmt.Println(banner())
	fmt.Printf("Making %d requests to %s with a concurrency level of %d.\n", f.n, f.url, f.c)
}
