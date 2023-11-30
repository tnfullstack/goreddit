package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	result := convert(5, 5.5)
	fmt.Println(result)

	str := "ABCDEFG"

	strSclice := strings.Split(str, "")
	fmt.Println(strSclice)

	for i, l := range str {
		fmt.Println(i, l, string(l))
	}

	toStdout(str)
}

func convert(i int, f float64) float64 {
	return float64(i) + f
}

func toStdout(s string) {
	_, err := io.WriteString(os.Stdout, s+"\n")
	if err != nil {
		log.Fatal(err)
	}
}
