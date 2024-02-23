package main

import (
	"fmt"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupValidAnagrams(strs))

}

func groupValidAnagrams(strs []string) [][]string {
	if len(strs) < 2 {
		return [][]string{strs}
	}
	m := make(map[[26]int][]string)
	for _, str := range strs {
		var count [26]int
		for _, c := range str {
			count[int(c)-97]++
		}
		m[count] = append(m[count], str)
	}
	aofa := make([][]string, len(m))
	i := 0
	for key := range m {
		var keySlice []string
		keySlice = m[key]
		aofa[i] = keySlice
		i++
	}

	return aofa
}
