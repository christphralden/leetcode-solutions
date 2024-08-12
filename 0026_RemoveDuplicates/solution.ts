// Time: 58ms (92.81%) | Memory: 52.88 (62.44%)
function removeDuplicates(nums: number[]): number {
  const set = new Set<number>(nums);
  nums.length = 0;
  nums.push(...set);
  return nums.length;
}
