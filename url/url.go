package url

// Refactoring the Parse Function

import (
	"errors"
	"fmt"
	"strings"
)

type URL struct {
	Scheme string
	Host   string
	Path   string
	// RawUrl string
}

// String method
func (u *URL) String() string {
	if u == nil {
		return ""
	}
	var s string
	if sc := u.Scheme; sc != "" {
		s += sc
		s += "://"
	}
	if h := u.Host; h != "" {
		s += h
	}
	if p := u.Path; p != "" {
		s += "/"
		s += p
	}
	return s
}

// testString
func (u *URL) testString() string {
	return fmt.Sprintf("scheme=%q, host=%q, path=%q", u.Scheme, u.Host, u.Path)
}

// GetHost returns u.Host, stripping port number if present.
func (u *URL) GetHost() string {
	host, _, ok := split(u.Host, ":", 0)
	if !ok {
		host = u.Host
	}
	return host
}

// Port returns the port part of u.Host, without the leading colon.
// If u.Host doesn't contain a port, Port returns an empty string.
func (u *URL) GetPort() string {
	_, port, _ := split(u.Host, ":", 0)
	return port
}

// Refactoring the Parse function
// ParseScheme parses rawurl into a URL structure.
func Parse(url string) (*URL, error) {
	// parse scheme
	scheme, rest, ok := parseScheme(url)
	if !ok {
		return nil, errors.New("missing scheme")
	}
	// parse hose
	host, path := parseHostPath(rest)

	// Store the correct scheme to URL struct
	return &URL{Scheme: scheme, Host: host, Path: path}, nil
}

// parseHost
func parseHostPath(str string) (host, path string) {
	host, path, ok := split(str, "/", 0)
	if !ok {
		host = str
	}
	return host, path
}

// parseScheme
func parseScheme(url string) (scheme, rest string, ok bool) {
	// get index position of "://"
	return split(url, "://", 1)
}

// Split s by sep.
// split returns empty string if it couldn't find sep in s at index n.
func split(s, sep string, n int) (a, b string, ok bool) {
	i := strings.Index(s, sep)
	if i < n {
		return "", "", false
	}
	return s[:i], s[i+len(sep):], true
}
