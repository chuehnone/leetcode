package llist

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetDecimalValue(head *ListNode) int {
	temp := head
	ans := 0
	for temp != nil {
		ans = ans << 1
		ans += temp.Val
		temp = temp.Next
	}

	return ans
}