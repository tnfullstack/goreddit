// tnutil package provice functions to do common tasks
package main

import (
	"fmt"
)

func main() {
	str1 := "ABCDEFG"
	str2 := byte(65)
	str3 := byte('A')
	str4 := rune(65)
	// str1 := []string{"A", "B", "C"}
	for _, r := range str1 {
		fmt.Printf("%c %b %x %x %X %o %d %q %T %#v\n", r, r, r, r, r, r, r, r, r, r)
	}
	fmt.Printf("%v, %v, %v, %v\n", str1, str2, str3, str4)

	p := &str1
	fmt.Printf("Address of %s is %p p = &str1 = %p\n", "str", &str1, p)
	fmt.Printf("Get the value stove in str1 = %s or %s\n", *p, *&str1)
}
