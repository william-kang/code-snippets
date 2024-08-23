class Solution {
  shortestDistance(words, word1, word2) {
    let i = words.indexOf(word1);
    let j = words.indexOf(word2);
    let d = words.length;

    while (i < words.length && j < words.length) {
      if (Math.abs(j - i) < d) d = Math.abs(j - i);
      if (i < j) {
        ++i;
        while (i < words.length && words[i] !== word1) ++i;
      } else {
        ++j
        while (j < words.length && words[j] !== word2) ++j;
      }
    }
    return d;
  }
}
