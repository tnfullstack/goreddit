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

func (u *URL) GetHost() string {
	return u.Host
}

func (u *URL) GetPort() string {
	port := ""
	if i := strings.Index(u.Host, ":"); i >= 0 {
		port = u.Host[i+1:]
	}
	return port
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

	// parse path
	host, path := rest, ""
	if i := strings.Index(rest, "/"); i >= 0 {
		host, path = rest[:i], rest[i+1:]
	}

	// Store the correct scheme to URL struct
	return &URL{Scheme: scheme, Host: host, Path: path}, nil
}
