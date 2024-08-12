// Time: 9ms (60.48%) | Memory: 6.53 (7.85%)
func trap(height []int) int {
  size := len(height)
  if(size == 0){
    return 0
  }
  left, right := 0, size-1
  leftMax, rightMax:= height[left], height[right]
  ans := 0 

  for left < right {
    if height[left] < height[right] {
      if height[left] > leftMax{
        leftMax = height[left]
      } else{
        ans += leftMax - height[left]
      }
      left++
    }else{
      if height[right] >= rightMax {
        rightMax = height[right]
      } else{
        ans += rightMax - height[right]
      }

      right--
    }
  }

  return ans
}
