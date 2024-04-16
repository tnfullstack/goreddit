package main

import "fmt"

// Modifying a slice with pointer

func reclassify(s *[]string) {
	*s = (*s)[0:8]
}

func main() {
	cities := []string{"Alantic City", "Benton", "Chicago", "Denver",
		"Fullton", "Greenville", "Huntsville", "Jasper", "Marion"}

	reclassify(&cities)
	fmt.Println(cities)
}
