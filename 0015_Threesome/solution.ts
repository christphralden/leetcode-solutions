// Time: 149ms (84.30%) | Memory: 65.24 (92.41%)
function threeSum(nums: number[]): number[][] {
  const length: number = nums.length;
  let ans: number[][] = [];

  nums.sort((a, b) => a - b);

  for (let i = 0; i < length; i++) {
    let left: number = i + 1;
    let right: number = length - 1;

    if (i > 0 && nums[i] == nums[i - 1]) continue;

    while (left < right) {
      const sum = nums[i] + nums[left] + nums[right];
      if (sum === 0) {
        ans.push([nums[i], nums[left], nums[right]]);

        while (left < right && nums[left] == nums[left + 1]) {
          left++;
        }

        while (left < right && nums[left] == nums[right]) {
          right--;
        }

        left++;
        right--;
        continue;
      }

      sum < 0 ? left++ : right--;
    }
  }
  return ans;
}
