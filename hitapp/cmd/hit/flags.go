// Extending the flag package

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/tvn9/gopl/url"
)

const usageText = `
Usage:
   hit [options] url
Options:`

// flags struct holds the flag fields
type flags struct {
	url  string
	n, c int
}

// number is a natrual number.
type number int

// toNumber is convenience function to convert p to *number.
func toNumber(p *int) *number {
	return (*number)(p)
}

func (n *number) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	switch {
	case err != nil:
		err = errors.New("parse error")
	case v <= 0:
		err = errors.New("should be positive")
	}
	*n = number(v)
	return err
}

func (n *number) String() string {
	return strconv.Itoa(int(*n))
}

// Using the flag variables
func (f *flags) parseFlags(s *flag.FlagSet, args []string) error {
	s.Usage = func() {
		fmt.Fprintln(s.Output(), usageText[1:])
		s.PrintDefaults()
	}

	// You no longer need to declare variables
	s.Var(toNumber(&f.n), "n", "Number of requests")
	s.Var(toNumber(&f.c), "c", "Concurrency level")
	if err := s.Parse(args); err != nil {
		return err
	}

	f.url = flag.Arg(0)

	if err := f.validate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		flag.Usage()
		return err
	}
	return nil
}

// validate post-conditions after parsing the flags
func (f *flags) validate() error {
	if f.c > f.n {
		return fmt.Errorf("-c=%d: should be less than or equal to -n=%d", f.c, f.n)
	}

	if f.c > runtime.NumCPU() {
		f.c = runtime.NumCPU()
		return fmt.Errorf("max available CPU is %d, option c range is 1 to 10", f.c)
	}

	if err := validateURL(f.url); err != nil {
		return fmt.Errorf("url: %w", err)
	}
	return nil
}

// validateURLFlag
func validateURL(s string) error {
	u, err := url.Parse(s)
	switch {
	case strings.TrimSpace(s) == "":
		err = errors.New("required")
	case err != nil:
		err = errors.New("parse error")
	case u.Scheme != "https" && u.Scheme != "http":
		err = errors.New("only supported scheme is http or https")
	case u.Host == "":
		err = errors.New("missing host")
	}
	return err
}

/*
// Using the standard library's parse package.
func (f *flags) parseFlags() error {
	var (
		u = flag.String("url", "", "HTTP server URL to (required) ")
		n = flag.Int("n", 20, "Number of requests")
		c = flag.Int("c", 5, "Concurrency level")
	)
	flag.Parse()
	f.url = *u
	f.n = *n
	f.c = *c
	return nil
}
*/

/*
// parseFunc is a command-line flag parser function.
type parseFunc func(string) error

// parseFalgs function parse os.Args[1:] command-line arguments manually
func (f *flags) parseFlags() (err error) {
	// a map of flag names and parsers function
	parsers := map[string]parseFunc{
		"url": f.urlVar(&f.url),
		"n":   f.intVar(&f.n),
		"c":   f.intVar(&f.c),
	}
	for _, arg := range os.Args[1:] {
		n, v, ok := strings.Cut(arg, "=")
		if !ok {
			continue // can't find the flag
		}
		pFunc, ok := parsers[strings.TrimPrefix(n, "-")]
		if !ok {
			continue // can't find a parse
		}
		if err = pFunc(v); err != nil {
			err = fmt.Errorf("invalid value %q for flag %s: %w", v, n, err)
			break // parsing error
		}
	}
	return err
}

// urlVar returns the function that call url.Parse then assign the result
// to the variable p (pointer to a string)
func (f *flags) urlVar(p *string) parseFunc {
	return func(s string) error {
		_, err := url.Parse(s)
		*p = s
		return err
	}
}

// intVar returns a function that take a string and perform calculation
// on the string then assign result to the variable on intVar function
// signature
func (f *flags) intVar(p *int) parseFunc {
	return func(s string) (err error) {
		*p, err = strconv.Atoi(s)
		return err
	}
}

*/
