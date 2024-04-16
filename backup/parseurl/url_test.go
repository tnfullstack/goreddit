package url

import (
	"fmt"
	"testing"
)

// Writing a failing test (url_test.go)
func TestParse(t *testing.T) {
	u := &URL{}

	// Test url
	const rawUrl = "https://foo.com/thanh"

	err := u.Parse(rawUrl)
	if err != nil {
		t.Errorf("fails to parse scheme, err %q\n", err)
	}

	fmt.Println(u.Scheme, u.Host, u.Path)
	want := "https"
	got := u.Scheme

	if want != got {
		t.Errorf("ParseScheme(%q).Scheme = %q; want %q\n", rawUrl, got, want)
	}

	want = "foo.com"
	got = u.Host

	if got != want {
		t.Errorf("Want %q but got %q\n", want, got)
	}

	want = "thanh"
	got = u.Path
	if got != want {
		t.Errorf("Want %q but got %q\n", want, got)
	}
}
