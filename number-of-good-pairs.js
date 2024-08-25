class Solution {
  numGoodPairs(nums) {
    let pairCount = 0;
    let numsCount = Array(101).fill(0);
    for (let num of nums) {
      numsCount[num] += 1;
    }
    for (let num of numsCount) {
      pairCount += (num - 1) / 2 * num;
    }
    return pairCount;
  }
}
