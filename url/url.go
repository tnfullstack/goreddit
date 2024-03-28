package url

import (
	"errors"
	"strings"
)

const correctScheme = "https"

type URL struct {
	Scheme string
	Host   string
	Path   string
	RawUrl string
}

// ParseScheme parses rawurl into a URL structure.
func (u *URL) ParseScheme(url string) error {
	if url == "" {
		return errors.New("url cannot be a blank string")
	}

	// store the raw url to URL struct
	u.RawUrl = url

	// get index position of "://"
	i := strings.Index(url, "://")
	if i < 0 {
		return errors.New("missing scheme")
	}

	tempScheme := url[:i]
	// validate scheme
	if tempScheme == correctScheme {
		// Store the correct scheme to URL struct
		u.Scheme = tempScheme
		return nil
	} else {
		u.Scheme = tempScheme
		return errors.New("invalid url scheme")
	}
}

// ParseHost parses the dns host name into a URL structure
func (u *URL) ParseHost(url string) error {

	return nil
}
