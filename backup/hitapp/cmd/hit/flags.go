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
	url       string
	n, c, rps int
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

// Using the flag variables
func (f *flags) parseFlags(s *flag.FlagSet, args []string) error {
	s.Usage = func() {
		fmt.Fprintln(s.Output(), usageText[1:])
		s.PrintDefaults()
	}

	// You no longer need to declare variables
	s.Var(toNumber(&f.n), "n", "Number of requests")
	s.Var(toNumber(&f.c), "c", "Concurrency level")
	s.Var(toNumber(&f.rps), "t", "Throttle request per second")
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
