package url_test

// This chapter covers

// Writing testable examples
// Producing executable documentation
// Measuring test coverage and benchmarking
// Refactoring the URL parser
// Differences between external and internal tests”

import (
	"fmt"
	"log"

	"github.com/tvn9/gopl/url"
)

// Adding additional example test
func ExampleURL_fields() {
	u, err := url.Parse("https://foo.com/go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u)
	// Output:
	// https
	// foo.com
	// go
	// https://foo.com/go
}

func ExampleURL() {
	u, err := url.Parse("https://foo.com/go")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	fmt.Println(u)
	// Output:
	// https://foo.com/go
}
