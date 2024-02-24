package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(validPalindrome("d  i<<<id"))

}

func validPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = regexp.MustCompile(`[^a-z0-9]`).ReplaceAllString(s, "")
	l := len(s)
	if l < 2 {
		return true
	}

	for c := range s {

		if s[c] != s[l-c-1] {
			return false
		} else if c > l-c-1 {
			return true
		}
	}
	return false
}
