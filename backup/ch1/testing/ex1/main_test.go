package main

import "testing"

func TestStrToNum(t *testing.T) {
	num, err := strToNum("1234567890")
	if err != nil {
		t.Fatal(err)
	}

	if num != 1234567890 {
		t.Fatal(err)
	}
}

func TestFailConv(t *testing.T) {
	_, err := strToNum("")
	if err == nil {
		t.Fatal(err)
	}
}
