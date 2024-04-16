package main

import "fmt"

func main() {
	// answer variable hold the address of that hold 42.
	answer := 42
	fmt.Println(answer, &answer, *&answer)

	// assign the address of 42 to address variable.
	address := &answer

	// print the address of 42, and dereferencing the address to access 42.
	fmt.Println(address, *address)

	// pointer type
	fmt.Printf("address is a type %T\n", address)
}
