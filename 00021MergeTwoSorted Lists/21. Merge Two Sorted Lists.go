package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	header := &ListNode{0, nil}
	p := header

	for l1 != nil || l2 != nil {
		if l1 == nil {
			p.Next = l2
			break
		}
		if l2 == nil {
			p.Next = l1
			break
		}

	}

	return header.Next
}

func main() {

}
