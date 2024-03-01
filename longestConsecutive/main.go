package main

func main() {

}

func longestConsecutive(nums []int) int {
    m := map[int]bool{}
    for _, val := range nums {
        if _, ok := m[val]; !ok {
            m[val] = false

        }
    }

    longest := 0

    for key := range m {
        count := 1
        deltaKey := key - 1
        _, ok := m[deltaKey]
        for ok {
            count++
            deltaKey--
            _, ok = m[deltaKey]
        }
        deltaKey = key + 1
        _, ok = m[deltaKey]
        for ok {
            count++
            deltaKey++
            _, ok = m[deltaKey]
        }
        if count > longest {
            longest = count
        }
        if longest > len(nums)/2 {
            return longest
        }
    }
    return longest
}
