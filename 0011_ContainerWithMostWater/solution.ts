function min(a: number, b: number) {
  return a < b ? a : b;
}
function maxArea(height: number[]): number {
  let left: number = 0;
  let right: number = height.length - 1;

  let ans: number = -999;
  let newSum: number = 0;

  while (left < right) {
    newSum = min(height[left], height[right]) * (right - left);
    if (newSum > ans) {
      ans = newSum;
    }

    height[left] < height[right] ? left++ : right--;
  }
  return ans;
}
