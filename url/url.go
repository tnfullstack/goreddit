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
		return &URL{}, errors.New("missing scheme")
	}

	// parse scheme
	scheme := url[:i]

	// Store the correct scheme to URL struct
	return &URL{Scheme: scheme}, nil
}
