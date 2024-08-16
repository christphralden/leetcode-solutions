// Time: 0ms (100%) | Memory: 2.54 (36.67%)
package main

import "fmt"

func max(arr *[]int, left int, right int) int {
  max := 0
  for i := left ; i < right ; i++ {
    if (*arr)[i] > (*arr)[max]{
      max = i
    }
  }
  return max
}

func swap(arr *[]int, x int, y int){
  (*arr)[x], (*arr)[y] = (*arr)[y], (*arr)[x]
}

func pancakeSort(arr []int) []int {
  ans := []int{}
  size := len(arr) - 1
  for i := size ; i > 0 ; i-- {
    maxIdx := max(&arr,0, i+1)

    left , right := 0, maxIdx
    ans = append(ans, maxIdx+1)

    for left < right {
      swap(&arr, left, right)
      left++
      right--
    }

    left , right = 0, i
    ans = append(ans, i+1)

    for left < right {
      swap(&arr, left, right)
      left++
      right--
    }
  }
  return ans 
}

func main(){
  tc1 := []int{3,2,4,1}
  fmt.Println(pancakeSort(tc1))
}
