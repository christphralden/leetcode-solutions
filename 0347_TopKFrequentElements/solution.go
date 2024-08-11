// Time: 6ms (93.93%) | Memory : 5.92 (91.47%)
func topKFrequent(nums []int, k int) []int {
  m :=  make(map[int]int)    

  for x := range nums{
    m[nums[x]] += 1
  }
  
  keys  := make([]int, 0, len(m))
  for key := range m{
    keys = append(keys, key)
  }

  sort.SliceStable(keys, func(i,j int) bool{
    return m[keys[i]] > m[keys[j]]
  })
  return keys[:k]
}
