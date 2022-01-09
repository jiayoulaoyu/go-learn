package main

import "fmt"

/**
两数相加
*/
func main() {
	l1 := initListNode([]int{5, 6})
	l2 := initListNode([]int{5, 4, 9})
	var res = addTwoNumbers(l1, l2)
	fmt.Println(res)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var nextVal int
	c := (l1.Val + l2.Val) % 10
	if l1.Val+l2.Val >= 10 {
		nextVal = 1
	}
	var res = &ListNode{Val: c}
	head := res
	tail := res
	var pre = l1.Next
	var next = l2.Next
	for (pre != nil) || (next != nil) {
		var currentNode ListNode
		preV := 0
		nextV := 0
		if pre != nil {
			preV = pre.Val
			pre = pre.Next
		}
		if next != nil {
			nextV = next.Val
			next = next.Next
		}
		sum := nextVal + preV + nextV
		if sum >= 10 {
			currentNode.Val = sum % 10
			nextVal = 1
		} else {
			currentNode.Val = sum
			nextVal = 0
		}
		tail.Next = &currentNode
		tail = tail.Next
	}
	if nextVal != 0 {
		var currentNode = &ListNode{Val: nextVal}
		tail.Next = currentNode
	}
	return head
}

func initListNode(arr []int) *ListNode {
	head := &ListNode{Val: arr[0]}
	tail := head
	for i := 1; i < len(arr); i++ {
		tail.Next = &ListNode{Val: arr[i]}
		tail = tail.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
