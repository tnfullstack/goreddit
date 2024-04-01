package url

// Tidying Up

// Writing tesable example
// Producing executable documentation
// Measuring test coverage and benchmarking
// Refacturing the URL parser
// Differences between external and internal tests

import (
	"fmt"
	"log"
	"testing"
)

// Share test TestURLPort test cases to test TestURLHostname
var hostTests = map[string]struct {
	in       string
	hostname string
	port     string
}{
	"with port":       {in: "foo.com:80", hostname: "foo.com", port: "80"},
	"with empty port": {in: "www.foo.com:", hostname: "www.foo.com", port: ""},
	"without port":    {in: "foo.com", hostname: "foo.com", port: ""},
	"ip with port":    {in: "127.0.0.1:90", hostname: "127.0.0.1", port: "90"},
	"ip without port": {in: "127.0.0.1", hostname: "127.0.0.1", port: ""},
	"messing scheme":  {in: "missingscheme.com", hostname: "missingscheme.com", port: ""},
}

// Adding test for covering edge-cases
func TestParseInvalidURLs(t *testing.T) {
	for name, tt := range hostTests {
		t.Run(fmt.Sprintf("%s/%s", name, tt.hostname), func(t *testing.T) {
			if _, err := Parse(tt.in); err == nil {
				t.Errorf("Parse(%q)=nil; want and error", tt.in)
			}
		})
	}
}

// Adding additional example test
func ExampleURL_fields() {
	u, err := Parse("https://foo.com/go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u)
	// Output:
	// https
	// foo.com
	// go
	// https://foo.com/go
}

// Testing the string method
func TestURLString(t *testing.T) {
	tests := map[string]struct {
		url  *URL
		want string
	}{
		"nil url":   {url: nil, want: ""},
		"empty url": {url: &URL{}, want: ""},
		"scheme":    {url: &URL{Scheme: "https"}, want: "https://"},
		"host":      {url: &URL{Scheme: "https", Host: "foo.com"}, want: "https://foo.com"},
		"path":      {url: &URL{Scheme: "https", Host: "foo.com", Path: "go"}, want: "https://foo.com/go"},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if g, w := tt.url, tt.want; g.String() != w {
				t.Errorf("url: %#v\ngot: %q\nwant: %q", g, g, w)
			}
		})
	}
}

// TestURLHostname
func TestURLHost(t *testing.T) {
	for name, tt := range hostTests {
		t.Run(fmt.Sprintf("%s/%s", name, tt.hostname), func(t *testing.T) {
			u := &URL{Host: tt.in}
			if got, want := u.GetHost(), tt.hostname; got != want {
				t.Errorf("for host %q; got %q; want %q", tt.in, got, want)
			}
		})
	}
}

// TestURLPort
func TestURLPort(t *testing.T) {
	for name, tt := range hostTests {
		t.Run(fmt.Sprintf("%s/%s", name, tt.in), func(t *testing.T) {
			u := &URL{Host: tt.in}
			if got, want := u.GetPort(), tt.port; got != want {
				t.Errorf("for host %q; got %q; want %q", tt.in, got, want)
			}
		})
		// ofcause subtest can be addted to with new t.Run to test other cases
		// t.Run(""), func(t *testing.T) { ... }
	}
}
