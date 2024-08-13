// Time: 51ms (93.41%) | Memory: 54.67 (5.49%)
function removeOuterParentheses(s: string): string {
  const stack: string[] = [];
  const result: string[] = [];

  for (const char of s) {
    if (char === "(") {
      if (stack.length > 0) {
        result.push(char);
      }
      stack.push(char);
    } else {
      stack.pop();
      if (stack.length > 0) {
        result.push(char);
      }
    }
  }

  return result.join("");
}

console.log(removeOuterParentheses("()()")); // Output: ""
console.log(removeOuterParentheses("(()())")); // Output: "()()"
console.log(removeOuterParentheses("(()())(())")); // Output: "()()()"
