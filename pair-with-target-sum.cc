using namespace std;

#include <iostream>
#include <vector>

class Solution {

  public:

  static vector<int> search(const vector<int> &arr, int targetSum) {
    int i = 0;
    int j = arr.size() - 1;

    while (i < j && arr[i] + arr[j] != targetSum) {
      if (arr[i] + arr[j] < targetSum) {
        ++i;
      } else {
        --j;
      }
    }

    if (i < j) return {i, j};
    else return {-1, -1};
  }

};
