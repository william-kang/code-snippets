class Solution {
  isAnagram(s, t) {
    if (s.length != t.length) return false;

    let arrS = new Array(26).fill(0);
    let arrT = new Array(26).fill(0);
    for (let i = 0; i < s.length; ++i) {
      ++arrS[s[i].charCodeAt(0) - 'a'.charCodeAt(0)];
      ++arrT[t[i].charCodeAt(0) - 'a'.charCodeAt(0)];
    }
    for (let i = 0; i < arrS.length; ++i) {
      if (arrS[i] != arrT[i]) return false;
    }
    return true;
  }
}
