package main

import (
	"fmt"
)

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(l *ListNode) *ListNode {
	h := &ListNode{Val: 2, Next: l}
	p1 := h
	p2 := p1.Next
	p3 := p2

	for p2 != nil {
		p3 = p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}
	h.Next.Next = nil
	// showListNodes(p1)
	l = p1
	return p1
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)

	over := false // 判断是否进位

	header := &ListNode{Val: 0, Next: nil}
	p := header

	for l1 != nil || l2 != nil {
		itemV := 0
		if l1 != nil {
			itemV = itemV + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			itemV = itemV + l2.Val
			l2 = l2.Next
		}
		if over {
			itemV = itemV + 1
			over = false
		}
		if itemV >= 10 {
			itemV = itemV - 10
			over = true
		}
		tempNode := &ListNode{Val: itemV, Next: nil}
		p.Next = tempNode
		p = tempNode
	}

	if over {
		tempNode := &ListNode{Val: 1, Next: nil}
		p.Next = tempNode
		p = tempNode
	}

	return header.Next
}

func generateListNodes(s []int) *ListNode {
	h1 := &ListNode{Val: 0, Next: nil}
	p := h1
	for _, nodeV := range s {
		tempNode := &ListNode{Val: nodeV, Next: nil}
		p.Next = tempNode
		p = tempNode
	}

	// h2 := ListNode{Val: 1, Next: nil}
	// p2 := h2
	// p3 := &p2
	// h3 := &h2
	// fmt.Println(&p2 == &h2, p3, h3)

	return h1.Next
}

func showListNodes(l1 *ListNode) {
	fmt.Println(l1)
	l2 := l1
	for l2 != nil {
		fmt.Print(l2.Val, " -> ")
		l2 = l2.Next
	}
	fmt.Println()
}

func main2() {
	l1Slice := []int{2, 4, 3}
	l1 := generateListNodes(l1Slice)

	l2Slice := []int{5, 6, 4}
	l2 := generateListNodes(l2Slice)

	// showListNodes(l1)

	// l1 = reverseList(l1)

	// showListNodes(l1)

	result := addTwoNumbers(l1, l2)
	showListNodes(result)

}
