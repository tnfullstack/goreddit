// url refacturing and save new code in parseurl
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

func (u *URL) Parse(url string) error {
	if ok := u.parseScheme(url); !ok {
		return errors.New("fails to parse scheme")
	}

	if ok := u.parseHost(url); !ok {
		return errors.New("fails to parse host")
	}

	if ok := u.parsePath(url); !ok {
		return errors.New("fails to parse path")
	}

	if ok := u.parseRawUrl(url); !ok {
		return errors.New("something wrong with URL string")
	}
	return nil
}

func (u *URL) parseScheme(url string) bool {
	i := strings.Index(url, "://")
	if i < 1 {
		return false
	}
	scheme := url[:i]
	u.Scheme = scheme
	return true
}

func (u *URL) parseHost(url string) bool {
	i := strings.Index(url, "://")
	if i < 1 {
		return false
	}
	rest := url[i+3:]
	host := rest
	if i := strings.Index(rest, "/"); i >= 0 {
		host = rest[:i]
	}
	u.Host = host
	return true
}

func (u *URL) parsePath(url string) bool {
	i := strings.Index(url, "://")
	if i < 1 {
		return false
	}
	rest := url[i+3:]
	path := ""
	if i := strings.Index(rest, "/"); i >= 0 {
		path = rest[i+1:]
	}
	u.Path = path
	return true
}

func (u *URL) parseRawUrl(url string) bool {
	if len(url) <= 0 {
		return false
	}
	u.RawUrl = url
	return true
}
