package main

import "strconv"

func main() {

}
func strToNum(str string) (num int64, err error) {
	num, err = strconv.ParseInt(str, 10, 64)
	return
}
