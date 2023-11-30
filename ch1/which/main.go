// Simple version of Unix which program in Go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: [command] [arguments] [Enter]")
		return
	}

	file := args[1]

	path := os.Getenv("PATH")
	filePath := filepath.SplitList(path)
	// fmt.Println(filePath)
	for _, dir := range filePath {
		fullPath := filepath.Join(dir, file)
		fmt.Println(fullPath)
		// Does not exist?
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			// Is it a regular file?
			if mode.IsRegular() {
				// Is it executable?
				if mode&0111 != 0 {
					fmt.Println(fullPath)
					return
				}
			}
		}
	}
}
