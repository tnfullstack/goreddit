package main

import (
	"flag"
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
