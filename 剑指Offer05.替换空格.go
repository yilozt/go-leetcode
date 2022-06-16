package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	b := strings.Builder{}

	for _, c := range s {
		if c == ' ' {
			b.WriteString("%20")
		} else {
			b.WriteRune(c)
		}
	}

	return b.String()
}

func assert(got string, want string) {
	if got != want {
		fmt.Println("err: got: '", got, "' want: '", want, "'")
	}
}

func main() {
	assert(replaceSpace("We are happy"), "We%20are%20happy")
	assert(replaceSpace(""), "")
	assert(replaceSpace(" We "), "%20We%20")
	assert(replaceSpace(" We are"), "%20We%20are")
	assert(replaceSpace("We are "), "We%20are%20")
	assert(replaceSpace("We are "), "We%20are%20")
}
