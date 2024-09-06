// TODO: Optimize use map
function lengthOfLongestSubstring(s: string): number {
  let i = 0;
  let maxLen = 0;
  let curr: string[] = [];
  const len = s.length;
  if (len == 1) return len;
  let j = 0;
  while (j < len) {
    if (!curr.includes(s.charAt(i))) {
      curr.push(s.charAt(i));
      i++;
      if (curr.length > maxLen) maxLen = curr.length;
    } else {
      curr.length = 0;
      j++;
      i = j;
    }
  }

  return maxLen;
}

console.log(lengthOfLongestSubstring("au"));
