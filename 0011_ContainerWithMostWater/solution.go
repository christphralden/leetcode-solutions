// Time: 53ms (93.94%) | Memory: 7.75 (68.98%)
func min(a *int, b *int) int{
  if(*a < *b){
    return *a  
  } else{
    return *b
  }
}
func maxArea(height []int) int {
  size := len(height)    

  left, right := 0, size-1
  ans := 0
  for left < right {
    
    if newSum := min(&height[left], &height[right]) * (right-left); newSum >= ans{
      ans = newSum
    }

    if height[left] < height[right]{
      left++
    } else{
      right--
    }
  }

  return ans
}
