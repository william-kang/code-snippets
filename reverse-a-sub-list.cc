using namespace std;

#include <iostream>

/*class ListNode {
public:
  int val = 0;
  ListNode *next;

  ListNode(int value) {
    this->val = value;
    next = nullptr;
  }
};*/

class Solution {
public:
  ListNode *reverse(ListNode *head, int p, int q) {
    ListNode *dum = new ListNode;
    dum->next = head;
    ListNode *beg = dum;
    for (int i = 0; i < p - 1; ++i) beg = beg->next;
    ListNode *end = head;
    for (int i = 1; i < q + 1; ++i) end = end->next;

    ListNode *a = beg->next;
    ListNode *b = beg->next->next;
    ListNode *c = beg->next->next ? beg->next->next->next : nullptr;

    a->next = end;
    while (b != end) {
      b->next = a;
      a = b;
      b = c;
      c = c ? c->next : nullptr;
    }
    beg->next = a;

    return dum->next;
  }
};
