// Boiling prints the boiling point of water.
package main

import "fmt"

const boilingF = 212.0

func main() {

	var f = boilingF
	var c = (f - 32) * 5 / 9

	fmt.Printf("Boiling point = %g℉ for %g℃\n", f, c) // Output boiling = 212℉ or 100℃
}
