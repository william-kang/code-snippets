using namespace std;

#include <iostream>

class ListNode {
public:
  int val = 0;
  ListNode *next;

  ListNode(int value) {
    this->val = value;
    next = nullptr;
  }
};

class Solution {
public:
  ListNode *reverse(ListNode *head) {
    if (!head || !head->next) return head;

    ListNode *a = head;
    ListNode *b = head->next;
    ListNode *c = head->next->next;

    a->next = nullptr;
    while (c) {
      b->next = a;
      a = b;
      b = c;
      c = c->next;
    }
    b->next = a;

    return b;
  }
};

int main() {
  ListNode *a = new ListNode(1);
  ListNode *b = new ListNode(2);
  ListNode *c = new ListNode(3);
  a->next = b;
  b->next = c;

  cout << a->val << ' ' << a->next->val << ' ' << a->next->next->val << endl;
  Solution s = Solution();
  a = s.reverse(a);
  cout << a->val << ' ' << a->next->val << ' ' << a->next->next->val << endl;
}
