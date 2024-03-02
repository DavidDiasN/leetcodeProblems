package main

func main() {

}

func lengthOfLongestSubstring(s string) int {
    m := map[byte]bool{}
    l := 0
    count := 0

    for r := range s{  
        _, still := m[s[r]]
        for  still {
            delete(m, s[l])
            l++
            _, still = m[s[r]]
        }
        m[s[r]] = false
        count = max(count, r-l + 1)
    }

    return count
}
