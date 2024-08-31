// ListNode struct represents a node in the linked list
// type ListNode struct {
//	Val int
//	Next  *ListNode
// }

// Solution struct is used to group the methods
type Solution struct{}

// hasCycle method checks for a cycle in the linked list
// It keeps the same logic and comments as the Java version
func (s *Solution) hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for i := 0; i < 100000; i++ {
		if fast == nil || fast.Next == nil { return false }
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow { return true }
	}
	return false
}
