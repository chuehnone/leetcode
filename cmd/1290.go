package main

import (
	"fmt"
	"leetcode/internal/llist"
)

// LeetCode: https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/
func main() {
	list := []int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}
	listLen := len(list)
	head := &llist.ListNode{}
	temp := head
	for i, v := range list {
		temp.Val = v
		if i < listLen-1 {
			temp.Next = &llist.ListNode{}
			temp = temp.Next
		}
	}

	fmt.Println(llist.GetDecimalValue(head))
}