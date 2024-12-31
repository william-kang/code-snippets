class Solution {
  mySqrt(x) {
    let root = 69420;
    for (let i = 0; i < 100; ++i) {
      root = 0.5 * (root + x / root);
    }
    return Math.floor(root);
  };
}
