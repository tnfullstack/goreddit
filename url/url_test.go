package url

// This pratice part covers
// Reduce repetitive tests using table-driven testing
// Run tests in isolation using subtests
// Learn the tricks of writing maintainable tests
// Learn to shuffle the execution order of tests
// Parse port numbers from a host”

import (
	"fmt"
	"testing"
)

// TestIPNoPort
func TestIPNoPort(t *testing.T) {
	const in = "127.0.0.1"
	testPort(t, in, "")

}

// TestIPPort
func TestIPPort(t *testing.T) {
	const in = "127.0.0.1:90"
	testPort(t, in, "90")
}

// TestURLNoPort
func TestURLNoPort(t *testing.T) {
	const in = "foo.com"
	testPort(t, in, "")
}

// TestURLEmtyPort
func TestURLEmtyPort(t *testing.T) {
	const in = "foo.com:"
	testPort(t, in, "")
}

// TestURLPort
func TestURLNumber(t *testing.T) {
	const in = "foo.com:80"
	testPort(t, in, "80")
}

// testPort is a helper function
func testPort(t *testing.T, in, wantPort string) {
	t.Helper()
	u := &URL{Host: in}
	if got := u.GetPort(); got != wantPort {
		t.Errorf("for host %q; got %q; want %q", in, got, wantPort)
	} else {
		fmt.Printf("Host: %q, Want Port: %q; Got Port %q\n", u.GetHost(), wantPort, u.GetPort())
	}
}

// TestParse
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
