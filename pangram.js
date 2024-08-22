class Solution {
  // Function to check if given sentence is pangram
  checkIfPangram(sentence) {
    sentence = sentence.toLowerCase();
    const set1 = new Set();
    for (let i = 0; i < sentence.length; ++i) {
      if (sentence.charCodeAt(i) >= 'a'.charCodeAt(0) &&
        sentence.charCodeAt(i) <= 'z'.charCodeAt(0)) {
        set1.add(sentence.charCodeAt(i) - 'a'.charCodeAt(0));
      }
    }
    return set1.size === 26;
  }
}
