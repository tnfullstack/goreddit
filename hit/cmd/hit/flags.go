package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// flags struct holds the flag fields
type flags struct {
	url  string
	n, c int
}

// parse
func (f *flags) parse() error {
	return nil
}

// parseFunc is a command-line flag parser function.
type parseFunc func(string) error

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
		parse, ok := parsers[strings.TrimPrefix(n, "-")]
		if !ok {
			continue // can't find a parse
		}
		if err = parse(v); err != nil {
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
