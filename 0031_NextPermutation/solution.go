// Time: 0ms (100%) | Memory: 2.60 (81.99%)
package main

import "fmt"

func findMaxFromEnd(nums []int) int{
  i := len(nums) - 1
  for i > 0 && nums[i] <= nums[i-1]{
    i--
  }
  return i
}

func findMaxRelative(nums []int, end int) int{
  i := len(nums) - 1
  for i > end-1 && nums[i] <= nums[end-1]{
    i--
  }
  return i
}

func reverseRelative(nums []int, from int){
  right := len(nums) - 1
  
  for right > from {
    swap(nums, right, from)
    from++
    right--
  }
}

func swap(nums []int, x int, y int){
  nums[x], nums[y] = nums[y], nums[x]
}

func nextPermutation(nums []int)  {
  maxIdx := findMaxFromEnd(nums)
  if maxIdx != 0 {
    swapper := findMaxRelative(nums, maxIdx)
    swap(nums, maxIdx-1, swapper)
  }
  reverseRelative(nums, maxIdx)
}

func main(){
  tc := []int{
    1,2,3,6,5,4,
  }
  tc2 := []int{
    // 1,2,3,
    1,1,5,
  }

  nextPermutation(tc)

  for i := 0 ; i < 6 ; i++{
    nextPermutation(tc2)
    fmt.Println(tc2)
  }
  fmt.Println(tc)
  fmt.Println(tc2)
}
