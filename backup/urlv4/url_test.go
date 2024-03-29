package url

// Adding test cases to test Parse host and path from rawurl

import (
	"fmt"
	"testing"
)

// Writing a failing test (url_test.go)
func TestParse(t *testing.T) {
	// Test url
	const rawurl = "https://foo.com/thanh"

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

	if got, want := u.Scheme, "https"; got != want {
		// t.Logf("ParseScheme(%q).Scheme = %q; want %q\n", rawurl, got, https)
		// t.FailNow()
		// t.Fatalf("ParseScheme(%q).Scheme = %q; want %q\n", rawurl, got, https)
		t.Errorf("ParseScheme(%q).Scheme = %q; want %q\n", rawurl, got, want)
	} else {
		fmt.Println("Got:", got, "Want:", want)
	}

	// Testing Parse host
	if got, want := u.Host, "foo.com"; got != want {
		t.Errorf("Parse(%q).Host = %q; want %q", rawurl, got, want)
	} else {
		fmt.Println("Got:", got, "Want:", want)
	}

	// Testing Parse path
	if got, want := u.Path, "thanh"; got != want {
		t.Errorf("Parse(%q).Host = %q; want %q", rawurl, got, want)
	} else {
		fmt.Println("Got:", got, "Want:", want)
	}
}
