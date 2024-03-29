package url

import (
	"testing"
)

// Writing a failing test (url_test.go)
func TestParse(t *testing.T) {
	// Test url
	const rawurl = "https://foo.com"

	u, err := Parse(rawurl)
	if err != nil {
		t.Logf("Parse(%q) err = %q, want nil", rawurl, err)
		t.Fail()
	}

	// Expected result
	https := "https"
	// http := "http"
	// Result received
	got := u.Scheme

	if got != https {
		t.Logf("ParseScheme(%q).Scheme = %q; want %q\n", rawurl, got, https)
		t.Fail()
	}
}
