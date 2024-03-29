package url

// Adding test cases to test Parse host from rawurl

import (
	"fmt"
	"testing"
)

// Writing a failing test (url_test.go)
func TestParse(t *testing.T) {
	// Test url
	const rawurl = "https://foo.com"

	u, err := Parse(rawurl)
	if err != nil {
		// t.Logf("Parse(%q) err = %q, want nil", rawurl, err)
		// t.Fail()
		// runtime.Goexit()
		// t.FailNow()
		// t.Fatalf = Logf + FailNow
		// t.Fatal(err) // replaced all method above with t.Fatalf
		t.Fatalf("Parse(%q) err = %q, want nil", rawurl, err)
		// t.Error = t.Log and t.Fail methods.
		// t.Errorf = t.Logf and t.Failf methods.
		t.Errorf("Parse(%q) err = %q, want nil", rawurl, err)

	}

	// Expected result
	https := "https"
	// http := "http"

	// Result received
	got := u.Scheme

	if got != https {
		// t.Logf("ParseScheme(%q).Scheme = %q; want %q\n", rawurl, got, https)
		// t.FailNow()
		// t.Fatalf("ParseScheme(%q).Scheme = %q; want %q\n", rawurl, got, https)
		t.Errorf("ParseScheme(%q).Scheme = %q; want %q\n", rawurl, got, https)
	}

	// Testing Parse host
	if got, want := u.Host, "foo.com"; got != want {
		t.Errorf("Parse(%q).Host = %q; want %q", rawurl, got, want)
	} else {
		fmt.Println("Got:", got, "Want:", want)
	}
}
