package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "/data00/new/eddy.txt"
	fmt.Println(basename(str))
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
