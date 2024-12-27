class Solution {
  containsDuplicate(nums) {
    var m = new Map();
    for (var i = 0; i < nums.length; ++i) {
      if (nums[i] in m) {
        return true;
      } else {
        m[nums[i]] = 1;
      }
    }
    return false;
  }
}
