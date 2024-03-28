package url

import "testing"

// Writing a failing test (url_test.go)
func TestParseScheme(t *testing.T) {
	u := &URL{}

	// Test url
	const rawUrl = "https://foo.com"

	err := u.ParseScheme(rawUrl)
	want := "https"
	got := u.Scheme
	if got != want {
		t.Fatalf("Parse(%q).Scheme = %q; want %q, Error: %q\n", rawUrl, got, want, err)
	}
}
