package main

import (
	"fmt"

	"github.com/tvn9/gopl/ch2/tempconv1/tempconv"
)

func main() {
	fmt.Printf("Brrrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println("Boiling C to F:", tempconv.CToF(tempconv.BoilingC))
	fmt.Printf("Kevin to Celsius: %.2f℃\n", tempconv.KToC(500))
}
