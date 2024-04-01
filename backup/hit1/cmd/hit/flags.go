package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type flags struct {
	url       string
	n, c, rps int
}

// number is a natural number.
type number int

// toNumber is a convenience function for converting p to *number.
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

func (f *flags) parse(s *flag.FlagSet, args []string) error {
	s.Usage = func() {
		fmt.Fprintln(os.Stderr, usageText[1:])
		flag.PrintDefaults()
	}
	// define flag variables
	s.Var(toNumber(&f.n), "n", "Number of request to make")
	s.Var(toNumber(&f.c), "c", "Concurrency level")
	s.Var(toNumber(&f.rps), "t", "Throttle requests per second")

	if err := s.Parse(args); err != nil {
		return err
	}
	if err := f.validate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		s.Usage()
		return err
	}

	f.url = flag.Arg(0)
	return nil
}

// validate post-conditions after parsing the flags.
func (f *flags) validate() error {
	if err := validateURL(f.url); err != nil {
		// return fmt.Errorf("invalid value %q for flag -url: %w", f.url, err)
		return fmt.Errorf("url: %w", err)
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
