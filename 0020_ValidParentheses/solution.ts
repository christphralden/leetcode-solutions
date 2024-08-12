function isValid(s: string): boolean {
  let splitted: string[] = s.split("");
  let stack: number[] = [];
  let rule = new Map<number, number>([
    [41, 40],
    [93, 91],
    [125, 123],
  ]);

  const keys = [...rule.keys()];
  for (let c of splitted) {
    const cInt: number = c.charCodeAt(0);
    if (keys.includes(cInt)) {
      if (stack[stack.length - 1] == rule.get(cInt)) {
        stack.pop();
      } else {
        return false;
      }
    } else {
      stack.push(cInt);
    }
  }
  return stack.length === 0;
}

console.log(isValid("]"));
