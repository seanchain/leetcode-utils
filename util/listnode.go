package util

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateNode(value int, next *ListNode) *ListNode {
	node := new(ListNode)
	node.Val = value
	node.Next = next
	return node
}

func (l *ListNode) Init(value int, next *ListNode) {
	l.Val = value
	l.Next = next
}

func CreateList(s []int) *ListNode {
	head := new(ListNode)
	cur := head
	for _, val := range s {
		node := CreateNode(val, nil)
		cur.Next = node
		cur = cur.Next
	}

	return head.Next
}

func CompareListNode(node1 *ListNode, node2 *ListNode) bool {
	for node1.Next != nil {
		var val1 = node1.Val
		if node2 == nil {
			return false
		}
		var val2 = node2.Val
		if val2 != val1 {
			return false
		}

		node1 = node1.Next
		node2 = node2.Next
	}
	return true
}

func PrintListNode(node *ListNode) {
	var printStr = ""
	for node.Next != nil {
		printStr += strconv.Itoa(node.Val)
		printStr += "->"
		node = node.Next
	}

	printStr += strconv.Itoa(node.Val)

	fmt.Println(printStr)
}
