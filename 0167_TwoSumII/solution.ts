// Time: 53ms (91.50%) | Memory: 52.34 (21.76%)
function twoSum(numbers: number[], target: number): number[] {
  let size: number = numbers.length;
  let left: number = 0;
  let right: number = size - 1;
  while (left < right) {
    let sum: number = numbers[left] + numbers[right];
    if (sum === target) {
      return [left + 1, right + 1];
    }
    sum < target ? left++ : right--;
  }
  return [];
}
