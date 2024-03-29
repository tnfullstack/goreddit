package url

import (
	"errors"
	"strings"
)

type URL struct {
	Scheme string
	Host   string
	Path   string
	RawUrl string
}

// ParseScheme parses rawurl into a URL structure.
func Parse(url string) (*URL, error) {
	// get index position of "://"
	i := strings.Index(url, "://")
	if i < 0 {
		return nil, errors.New("missing scheme")
	}

	// parse scheme, and host
	scheme, rest := url[:i], url[i+3:]

	// Store the correct scheme to URL struct
	return &URL{Scheme: scheme, Host: rest}, nil
}
