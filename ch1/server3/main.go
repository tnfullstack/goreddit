// Server3 is a minimal "echo" and counter server, with added functionality
// to handler to report on headers and form data that it receives, making
// the sever useful for inspecting and debugging requests:
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	repeat := 30
	fmt.Fprintln(w, strings.Repeat("-", repeat)) // ---------------------------------

	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	fmt.Fprintln(w, strings.Repeat("-", repeat))

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintln(w, strings.Repeat("-", repeat))

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintln(w, strings.Repeat("-", repeat))

	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	fmt.Fprintln(w, strings.Repeat("-", repeat))

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
