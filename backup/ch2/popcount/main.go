package main

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

/*
func PopCount(x uint64) int {
	var popCount int
	for i := 0; i < int(x); i++ {
		p := int(pc[byte(x>>(i*8))])
		popCount += p
	}
	return popCount
}
*/

func main() {

	pop := PopCount(9)
	fmt.Printf("Pop count = %d\n", pop)
}
