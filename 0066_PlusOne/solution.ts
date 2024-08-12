// Time: 51ms (85.37%) | Memory: 51.24 (65.37%)
function plusOne(digits: number[]): number[] {
  const len = digits.length;
  let i = 1;
  digits[len - i]++;

  while (true) {
    if (digits[len - i] == 10) {
      digits[len - i] = 0;
      if (digits[len - i - 1] === undefined) {
        digits.unshift(1);
      } else {
        digits[len - i - 1]++;
      }
      i++;
    } else {
      return digits;
    }
  }
}

// console.log(plusOne([8, 9, 9, 9]));
// console.log(plusOne([9, 9]));
