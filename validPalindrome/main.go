package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	//	fmt.Println(validPalindrome("d  i<<<id"))
	fmt.Println(int('A'))
	fmt.Println(int('Z'))
	fmt.Println(int('a'))
	fmt.Println(int('z'))
	fmt.Println(int('0'))
	fmt.Println(int('9'))

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

func neetValidPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1

	for l < r {
		if !isAlNum(int(s[l])) {
			l++
			continue
		}
		if !isAlNum(int(s[r])) {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func isAlNum(c int) bool {
	return (c >= 65 && c <= 90) || (c >= 97 && c <= 122) || (c >= 48 && c <= 57)

}
