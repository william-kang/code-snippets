#include <iostream>
#include <cmath>

class Solution {
public:
  int mySqrt(int x) {
    if (x < 2) return x;

    int left = 2;
    int right = std::floor(x / 2);

    while (left <= right) {
      long num = std::floor((left + right) / 2);
      if (num * num > x) right = num - 1;
      else if (num * num < x) left = num + 1;
      else return num;
    }

    return right;
  }
};
