package url

const correctScheme = "https"

type URL struct {
	Scheme string
	Host   string
	Path   string
	RawUrl string
}

// ParseScheme parses rawurl into a URL structure.
func Parse(url string) error {

	// 	// // get index position of "://"
	// 	// i := strings.Index(url, "://")
	// 	// if i < 0 {
	// 	// 	return errors.New("missing scheme")
	// 	// }

	// 	// // parse scheme
	// 	// scheme, rest := url[:i], url[i+3:]

	// 	// // parse host, and path
	// 	// host, path := rest, ""
	// 	// if i := strings.Index(rest, "/"); i >= 0 {
	// 	// 	host, path = rest[:i], rest[i+1:]
	// 	// }

	// 	// Store the correct scheme to URL struct
	return nil
}
