package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

type flags struct {
	url  string
	n, c int
}

/*
func (f *flags) parse() (err error) {
	for _, arg := range os.Args[1:] {
		n, v, ok := strings.Cut(arg, "=")
		if !ok {
			continue // can't find a parser
		}
		parse, ok := parsers[strings.TrimPrefix(n, "-")]
		if !ok {
			continue // can't find a parser
		}
		if err = parse(v); err != nil {
			err = fmt.Errorf("invalid value %q for flag %s: %w", v, n, err)
			break // parsing error
		}
	}
	return err
}
*/

func (f *flags) parse() error {
	// define flag variables
	flag.StringVar(&f.url, "url", "", "HTTP server `URL` (required)")
	flag.IntVar(&f.n, "n", f.n, "Number of request")
	flag.IntVar(&f.c, "c", f.c, "Concurrency level")
	flag.Parse()

	if err := f.validate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		flag.Usage()
		return err
	}

	// var (
	// 	u = flag.String("url", "", "HTTP server `URL` (required)")
	// 	n = flag.Int("n", f.n, "Number of requests")
	// 	c = flag.Int("c", f.c, "Concurrency level")
	// )
	// flag.Parse()
	// f.url = *u
	// f.n = *n
	// f.c = *c
	return nil
}

// validate post-conditions after parsing the flags.
func (f *flags) validate() error {
	if err := validateURL(f.url); err != nil {
		return fmt.Errorf("invalid value %q for flag -url: %w", f.url, err)
	}
	if f.c > f.n {
		return fmt.Errorf("-c=%d: should be less than or equal to -n=%d", f.c, f.n)
	}
	return nil
}

func validateURL(s string) error {
	if strings.TrimSpace(s) == "" {
		return errors.New("required")
	}
	u, err := url.Parse(s)
	switch {
	case strings.TrimSpace(s) == "":
		err = errors.New("required")
	case err != nil:
		err = errors.New("parse error")
	case u.Scheme != "http":
		err = errors.New("only supported scheme is http")
	case u.Host == "":
		err = errors.New("missing host")
	}
	return err
}
