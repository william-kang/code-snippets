class Solution {
  // Function to check if given sentence is pangram
  checkIfPangram(sentence) {
    sentence = sentence.toLowerCase();
    const s = new Set();
    for (let i = 0; i < sentence.length; ++i) {
      if (sentence[i] >= 'a' && sentence[i] <= 'z' && !s.has(sentence[i])) {
        s.add(sentence[i]);
      }
    }
    return s.size == 26;
  }
}
