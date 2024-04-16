// Find min and max values and discard invalid values
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: [command] [1 2 3 4] [Enter]")
		return
	}

	fmt.Println("User input:", args)
	var min, max float64
	for i := 1; i < len(args); i++ {
		n, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			continue
		}
		if i == 1 {
			min = n
			max = n
			continue
		}
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
