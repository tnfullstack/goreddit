package main

import (
	"fmt"
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
	fmt.Println(banner())
	fmt.Println(usage())
}
