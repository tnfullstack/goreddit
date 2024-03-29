package url

import (
	"testing"
)

// Writing a failing test (url_test.go)
func TestParse(t *testing.T) {

	// Test url
	const rawurl = "broken url"

	err := Parse(rawurl)
	if err != nil {
		// t.Log(err)
		t.Logf("Parse(%q) err = %q, want nil", rawurl, err)
		t.Fail()
	}

	// want := correctScheme
	// got := u.Scheme

	// if want != got {
	// 	t.Log("test fail")
	// 	t.Logf("ParseScheme(%q).Scheme = %q; want %q\n", rawUrl, got, want)

	// 	t.Fail()
	// }

}
