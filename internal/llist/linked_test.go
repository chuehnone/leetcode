package llist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDecimalValue(t *testing.T) {
	head := &ListNode{}
	list := []int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}
	listLen := len(list)
	temp := head
	for i, v := range list {
		temp.Val = v
		if i < listLen-1 {
			temp.Next = &ListNode{}
			temp = temp.Next
		}
	}

	assert.Equal(t, 18880, GetDecimalValue(head))
}
