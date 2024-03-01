package main

func main() {

}

func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)
    arr := make([][]int, len(nums))

    for _, val := range nums {
        m[val]++
    }
    for key := range m {
        arr[m[key]-1] = append(arr[m[key]-1], key)
    }

    res := []int{}

    i := len(nums) - 1
    for i >=0{
        for _, val := range arr[i] {
            res = append(res, val)
            if len(res) == k {
                return res
            }
        }
        i--
    }
    
    return res
}


