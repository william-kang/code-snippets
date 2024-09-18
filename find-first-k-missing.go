import java.util.*;

class Solution {
  public static List<Integer> findNumbers(int[] nums, int k) {
    int bigNums[] = new int[1000];
    for (int i = 0; i < nums.length; i++) {
      bigNums[i] = nums[i];
    }
    for (int i = 1; i <= bigNums.length; i++) {
      while (
        bigNums[i - 1] > 0 && bigNums[i - 1] <= bigNums.length &&
        bigNums[i - 1] != i && bigNums[i - 1] != bigNums[bigNums[i - 1] - 1]
      ) {
        int temp = bigNums[i - 1];
        bigNums[i - 1] = bigNums[temp - 1];
        bigNums[temp - 1] = temp;
      }
    }
    List<Integer> missingNumbers = new ArrayList<>();
    for (int i = 1; i <= bigNums.length && missingNumbers.size() < k; i++) {
      if (bigNums[i - 1] != i) missingNumbers.add(i);
    }
    for (int i = 1001; missingNumbers.size() < k; missingNumbers.add(i++)) {}
    return missingNumbers;
  }
}
