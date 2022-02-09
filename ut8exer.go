package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	name := "hihoheyhu"

	for len(name) > 0 {
		r, size := utf8.DecodeLastRuneInString(name)

		fmt.Printf("%c %v\n", r, size)

		name = name[:len(name)-size]
	}
}
