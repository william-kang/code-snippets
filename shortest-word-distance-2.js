class Solution {
  shortestDistance(words, word1, word2) {
    let i = 0, j = 0, d = words.length;
    while (words[i] != word1) ++i;
    while (words[j] != word2) ++j;
    d = Math.abs(i - j);

    while (true) {
      if (i < j) {
        ++i;
        while (i < words.length && words[i] != word1) ++i;
      } else {
        ++j;
        while (j < words.length && words[j] != word2) ++j;
      }
      if (i < words.length && j < words.length) d = Math.min(Math.abs(i - j), d);
      else return d;
    }
  }
}
