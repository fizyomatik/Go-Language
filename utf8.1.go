package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	name := "inanä    "
	name = strings.TrimRight(name, "")
	l := utf8.RuneCountInString(name)
	fmt.Println(l)
	fmt.Println(len(name))

}
