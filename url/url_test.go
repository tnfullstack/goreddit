package url

import "testing"

// Writing a failing test (url_test.go)
func TestParseScheme(t *testing.T) {
	u := &URL{}

	// Test url
	const rawUrl = "https://foo.com"

	err := u.ParseScheme(rawUrl)
	if err != nil {
		t.Logf("fails to parse scheme, err %q\n", err)
	}

	want := correctScheme
	got := u.Scheme

	if want != got {
		t.Fatalf("ParseScheme(%q).Scheme = %q; want %q\n", rawUrl, got, want)
	}
}
