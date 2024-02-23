package main

import "fmt"

func main() {
	boo := validAnagram("string", "string")
	fmt.Println(boo)

}

func validAnagram(s, t string) bool {
	m := make(map[byte]int)

	if len(s) != len(t) {
		return false
	}
	for i := range s {
		m[s[i]]++

		m[t[i]]--
	}

	for key := range m {
		if m[key] != 0 {
			return false
		}
	}
	return true
}
