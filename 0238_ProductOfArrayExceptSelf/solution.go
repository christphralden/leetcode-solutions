// Time: 28ms (97.34%) | Memory: 8.51 (49.62%)
func productExceptSelf(nums []int) []int {
    length := len(nums)
    
    ans := make([]int, length)
    
    ans[0] = 1
    
    for i := 1; i < length; i++ {
        ans[i] = ans[i-1] * nums[i-1]
    }

    suf := 1
    for i := length - 2; i >= 0; i-- {
        suf *= nums[i+1]
        ans[i] *= suf
    }
    
    return ans
}
