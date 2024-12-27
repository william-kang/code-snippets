class Solution {
  isVowel(c) {
    return c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' ||
      c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u';
  }
  reverseVowels(s) {
    let a = s.split('');
    let i = 0, j = a.length - 1;

    while (true) {
      while (i < a.length && !this.isVowel(a[i])) ++i;
      while (j > 0 && !this.isVowel(a[j])) --j;
      if (i < j) {
        let temp = a[i];
        a[i] = a[j];
        a[j] = temp;
        ++i;
        --j;
      } else break;
    }

    return a.join('');
  }
}
