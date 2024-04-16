package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"github.com/tvn9/gopl/hitapp/hit"
)

const (
	bannerText = `
 __  __     __     ______
/\ \_\ \   /\ \   /\__  _\
\ \  __ \  \ \ \  \/_/\ \/ 
 \ \_\ \_\  \ \_\    \ \_\
  \/_/\/_/   \/_/     \/_/
` // cannot remove first newline character with constant [1:]
)

// Refactor the bannerText and usageText with functions
// This will solve the slicing problem when using const
func banner() string { return bannerText[1:] }

// func usage() string { return usageText[1:] }

// run
func run(s *flag.FlagSet, args []string, out io.Writer) error {
	f := &flags{
		n: 100,
		c: runtime.NumCPU(),
	}
	if err := f.parseFlags(s, args); err != nil {
		return err
	}

	fmt.Fprintln(out, banner())
	fmt.Fprintf(out, "Making %d requests to %s with a concurrency level %d\n", f.n, f.url, f.c)

	if f.rps > 0 {
		fmt.Fprintf(out, "(RPS: %d)\n", f.rps)
	}

	request, err := http.NewRequest(http.MethodGet, f.url, http.NoBody)
	if err != nil {
		return err
	}
	c := &hit.Client{
		RPS: f.rps,
	}
	sum := c.Do(request, f.n)
	sum.Fprint(out)

	return nil
}

func main() {
	if err := run(flag.CommandLine, os.Args[1:], os.Stdout); err != nil {
		os.Exit(1)
	}
}
