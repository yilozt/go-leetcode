package main

import "strings"

func replaceSpace(s string) string {
	var tmp strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			tmp.WriteString("%20")
		} else {
			tmp.WriteByte(s[i])
		}
	}
	return tmp.String()
}
