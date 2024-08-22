class Solution {
  isPalindrome(s) {
    s = s.toLowerCase().replace(/\W/g, '');
    let i = 0, j = s.length - 1;
    while (i < j) {
      if (s[i] !== s[j]) return false;
      i++;
      j--;
    }
    return true;
  }
}
