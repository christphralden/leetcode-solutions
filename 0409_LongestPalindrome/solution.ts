// Time: 52ms (92.91%) | Memory: 52.63 (16.89%)
function longestPalindrome(s: string): number {
  let map = new Map();

  for (let x of s) {
    map.set(x, (map.get(x) ?? 0) + 1);
  }
  let isOne = 0;
  let count = 0;
  for (let x of map.entries()) {
    if ((x[1] == 1 || x[1] % 2 == 1) && !isOne) {
      isOne = 1;
    }
    count += 2 * Number(Math.floor(x[1] / 2));
  }
  return count + isOne;
}

console.log(longestPalindrome("abcccdd"));
