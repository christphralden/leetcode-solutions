// Time: 63ms (89.21%)  | Memory: 53.24 (88.40%)
function topKFrequent(nums: number[], k: number): number[] {
  if (nums.length === k) return nums;

  let map: Map<number, number> = new Map<number, number>();
  const len: number = nums.length;

  for (let i: number = 0; i < len; i++) {
    map.set(nums[i], (map.get(nums[i]) ?? 0) + 1);
  }
  const mapDesc: Map<number, number> = new Map(
    [...map.entries()].sort((a, b) => b[1] - a[1]),
  );
  return [...mapDesc.keys()].slice(0, k);
}
