// Time: 123ms (96.43%) | Memory: 57.78 (34.52%)
function numRescueBoats(people: number[], limit: number): number {
  let left = 0;
  let right = people.length - 1;
  let ans = 0;

  people.sort((a, b) => a - b);
  while (left <= right) {
    if (people[left] + people[right] <= limit && right != left) {
      left++;
    }
    ans++;
    right--;
  }

  return ans;
}

console.log(
  numRescueBoats(
    [2695, 22727, 26302, 27700, 25273, 26944, 27691, 16217, 11739, 21158],
    29998,
  ),
);
