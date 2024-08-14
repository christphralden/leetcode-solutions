// Time: 158ms (100%) | Memory: 63.06 (27.78%)
function numSubseq(nums: number[], target: number): number {
  const mod = 1e9 + 7;
  const len = nums.length;
  let right = len - 1;
  let left = 0;
  let ans = 0;

  nums.sort((a, b) => a - b);

  let powMemo: number[] = [1];

  for (let i = 1; i < len; i++) {
    powMemo[i] = (powMemo[i - 1] * 2) % mod;
  }

  while (left <= right) {
    if (nums[left] + nums[right] <= target) {
      ans = (ans + powMemo[right - left]) % mod;
      left++;
    } else {
      right--;
    }
  }
  return ans;
}
console.log(numSubseq([2, 3, 3, 4, 6, 7], 12));
