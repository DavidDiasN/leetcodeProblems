package main

import ("fmt" "sort")

func main() {
  fmt.Println("yee")
}

func threeSum(nums []int) [][]int {
  var result [][]int
  lastPoint := 0
  m := make(map[int]int)
  
  for i := range nums {
    if _, ok := m[nums[i]]; !ok {
      m[nums[i]] = i
    } 

  }
    
}
