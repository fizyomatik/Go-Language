package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("EMö")
	fmt.Println(utf8.Valid(b))

}
