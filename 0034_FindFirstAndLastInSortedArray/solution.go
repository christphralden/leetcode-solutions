package main

import (
  "math"
)

func search(nums []int, low int, high int, x int) int{
  for low <= high {
    mid := int(math.Floor(float64((high/2))))

    if nums[mid] == x {
      return mid
    }
    if nums[mid] < x {
      low = mid + 1 
    } else {
      high = mid - 1
    }
  }

  return -1
}
func searchRange(nums []int, target int) []int {
  left := 0
  right := len(nums) - 1

  search(nums, left, right, target) 
   
  
  return nums
}

func main(){
  
  tc1 := []int{5,7,7,8,8,10}
  searchRange(tc1, 8)
}
