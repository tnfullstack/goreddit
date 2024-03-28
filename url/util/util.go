package util

func GetStrIndex(str string, p rune) (index int) {
	if str == "" {
		index = -1
		return index
	}

	for i, e := range str {
		if e == p {
			index = i
		} else {
			index = -1
		}
	}
	return index
}
