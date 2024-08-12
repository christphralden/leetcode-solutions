// Time: 39ms (99.51%) | Memory: 51.12 (72.68%)
function lengthOfLastWord(s: string): number {
  return s.trim().split(" ").pop()?.length ?? 0;
}
