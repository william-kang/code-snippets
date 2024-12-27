class Solution {
  isPalindrome(s) {
    s = s.toLowerCase();
    let a = [];
    for (let i = 0; i < s.length; ++i) {
      if (s[i] >= 'a' && s[i] <= 'z' || s[i] >= '0' && s[i] <= '9') a.push(s[i]);
    }
    let i = 0, j = a.length - 1;
    while (i < j) {
      if (a[i] != a[j]) return false;
      ++i;
      --j;
    }
    return true;
  }
}
