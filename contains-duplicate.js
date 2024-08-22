class Solution {
  containsDuplicate(nums) {
    const map = new Map();

    for (var i = 0; i < nums.length; ++i) {
      if (map.has(nums[i])) {
        return true;
      } else {
        map.set(nums[i], 0);
      }
    }

    return false;
  }
}
