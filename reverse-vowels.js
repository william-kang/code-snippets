class Solution {
  reverseVowels(s) {
    let arr = s.split("");
    let i = 0, j = arr.length - 1;
    while (true) {
      while (i < arr.length && !this.isVowel(arr[i])) ++i;
      while (j >= 0 && !this.isVowel(arr[j])) --j;
      if (i < arr.length && j >= 0 && i < j) {
        [arr[i], arr[j]] = [arr[j], arr[i]];
        console.log('swapped', i, j);
        ++i;
        --j;
      } else {
        break;
      }
    }
    return arr.join("");
  }

  isVowel(c) {
    return 'aeiouAEIOU'.includes(c);
  }
}
