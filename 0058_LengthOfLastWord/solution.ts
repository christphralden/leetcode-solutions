// Time: O(n) | Memory: O(n)
function lengthOfLastWord(s: string): number {
  return s.trim().split(" ").pop()?.length ?? 0;
}

console.log(lengthOfLastWord("   fly me   to   the moon  "));
