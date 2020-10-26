package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}

	first := &ListNode{0, head}
	second := first

	for i := 0; i <= n; i++ {
		first = first.Next
	}

	for first.Next != nil {
		first = first.Next
		second = second.Next
	}

	temp := second.Next
	second.Next = temp.Next

	return head
}

func main() {

}
