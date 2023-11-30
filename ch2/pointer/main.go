// pointer demonstrates how pointer work in Go
package main

import "fmt"

func main() {
	num := 42

	// get the address of num veriable
	fmt.Printf("Address of %s is: %v\n", "num", &num)

	// addrOfnum holds the address of veriable num which point to value 12123
	addrOfnum := &num
	fmt.Printf("Get value of num using its address %v\n", *addrOfnum)

	// modify the value from num veriable
	*addrOfnum = *addrOfnum - 2
	fmt.Printf("num after subtract 2 = %d\n", num)

	// increase add 1 to the passing value
	result := increase(&num)
	fmt.Printf("num after add 1 = %d\n", result)
	fmt.Printf("num after add 1 = %d\n", num)
}

// increase add 1 to the passing value
func increase(n *int) int {
	*n++
	return *n
}
