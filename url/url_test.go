package url

// This pratice part covers
// Reduce repetitive tests using table-driven testing
// Run tests in isolation using subtests
// Learn the tricks of writing maintainable tests
// Learn to shuffle the execution order of tests
// Parse port numbers from a host”

// In this version - Writing test with subtests

import (
	"fmt"
	"testing"
)

// Writing subtest
func TestURLPort1(t *testing.T) {
	t.Run("with port", func(t *testing.T) {
		const in = "foo.com:80"
		u := &URL{Host: in}
		if got, want := u.GetPort(), "800"; got != want {
			t.Errorf("for host %q; got %q; want %q", in, got, want)
		}
	})
	t.Run("host empty port", func(t *testing.T) {
		const in = "foo.com:"
		u := &URL{Host: in}
		if got, want := u.GetPort(), ""; got != want {
			t.Errorf("for host %q; got %q; want %q", in, got, want)
		}
	})
	t.Run("host no port", func(t *testing.T) {
		const in = "foo.com"
		u := &URL{Host: in}
		if got, want := u.GetPort(), ""; got != want {
			t.Errorf("for host %q; got %q; want %q", in, got, want)
		}
	})
	t.Run("ip with port", func(t *testing.T) {
		const in = "127.0.0.1:90"
		u := &URL{Host: in}
		if got, want := u.GetPort(), "900"; got != want {
			t.Errorf("for host %q; got %q; want %q", in, got, want)
		}
	})
	t.Run("IP no port", func(t *testing.T) {
		const in = "127.0.0.1"
		u := &URL{Host: in}
		if got, want := u.GetPort(), ""; got != want {
			t.Errorf("for host %q; got %q; want %q", in, got, want)
		}
	})
}

// Writing test with table testing

// TestURLPort
// func TestURLPort(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		in       string
// 		expected string
// 	}{
// 		// first test save in index 1
// 		{name: "host with port", in: "foo.com:80", expected: "80"},
// 		{name: "host empty port", in: "foo.com:", expected: ""},
// 		{name: "host no port", in: "foo.com", expected: ""},
// 		{name: "IP with port", in: "127.0.0.1:90", expected: "90"},
// 		{name: "IP no port", in: "127.0.0.1", expected: ""},
// 	}

// 	str := fmt.Sprintf("%-3s %-26s %-14s %-6s %s", strings.ToUpper("No"), "Test Name", "Host Name", "Got", "Expected")
// 	fmt.Println(strings.ToUpper(str))

// 	for i, tt := range tests {
// 		u := &URL{Host: tt.in}
// 		if got, expected := u.GetPort(), tt.expected; got != expected {
// 			t.Errorf("%-3d %-26s %-14s %-6q %q\n", i+1, tt.name, tt.in, got, expected)
// 		} else {
// 			fmt.Printf("%-3d %-26s %-14s %-6q %q\n", i+1, tt.name, tt.in, got, expected)
// 		}
// 	}
// }

// Writing individual test case with testPort Helper funciton

// TestIPNoPort
// func TestIPNoPort(t *testing.T) {
// 	const in = "127.0.0.1"
// 	testPort(t, in, "")
// }

// // TestIPPort
// func TestIPPort(t *testing.T) {
// 	const in = "127.0.0.1:90"
// 	testPort(t, in, "90")
// }

// // TestURLNoPort
// func TestURLNoPort(t *testing.T) {
// 	const in = "foo.com"
// 	testPort(t, in, "")
// }

// // TestURLEmtyPort
// func TestURLEmtyPort(t *testing.T) {
// 	const in = "foo.com:"
// 	testPort(t, in, "")
// }

// // TestURLPort
// func TestURLNumber(t *testing.T) {
// 	const in = "foo.com:80"
// 	testPort(t, in, "80")
// }

// // testPort is a helper function
// func testPort(t *testing.T, in, wantPort string) {
// 	t.Helper()
// 	u := &URL{Host: in}
// 	if got := u.GetPort(); got != wantPort {
// 		t.Errorf("for host %q; got %q; want %q", in, got, wantPort)
// 	} else {
// 		fmt.Printf("Host: %q, Want Port: %q; Got Port %q\n", u.GetHost(), wantPort, u.GetPort())
// 	}
// }

// Writing test with one function and if statement to check differnt test scenarios

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
