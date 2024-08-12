// Time: 37ms (92.18%) | Memory: 7.70 (59.61%)
func process(nums *[]int, ans *[][]int){
    n := len(*nums)

    left, right := 0, n-1

    for i := 0; i < n-2; i++ {
        if i > 0 && (*nums)[i] == (*nums)[i-1] {
            continue
        }
        left = i+1
        right = n-1
        for left < right {
            sum := (*nums)[i] + (*nums)[left] + (*nums)[right]

            if sum == 0 {
                *ans = append(*ans, []int{(*nums)[i], (*nums)[left], (*nums)[right]})

                for left < right && (*nums)[left] == (*nums)[left+1] {
                    left++
                }
                for left < right && (*nums)[right] == (*nums)[right-1] {
                    right--
                }

                left++
                right--
            } else if sum < 0 {
                left++
            } else {
                right--
            }
        }
    }


}
func threeSum(nums []int) [][]int {
    ans := [][]int{}
    sort.Ints(nums)
    process(&nums, &ans)
    return ans
}
