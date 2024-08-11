function containsDuplicate(nums: number[]): boolean {
  const arr: number[] = [];
  let i = 0;
  const size = nums.length;
  while (i < size) {
    if ((arr[nums[i]] = (arr[nums[i]] ?? 0) + 1) > 1) return true;
    i++;
  }
  return false;
}
