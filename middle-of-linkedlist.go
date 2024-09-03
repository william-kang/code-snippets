import ()

// type ListNode struct {
// 	Val int
// 	Next  *ListNode
// }

type Solution struct{}

// findMiddle finds the middle node in a linked list
func (s *Solution) findMiddle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
