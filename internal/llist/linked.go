package llist

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetDecimalValue(head *ListNode) int {
	temp := head
	ans := 0
	for temp != nil {
		ans += temp.Val
		temp = temp.Next
		if temp != nil {
			ans *= 2
		}
	}

	return ans
}