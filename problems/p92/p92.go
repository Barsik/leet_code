package p92

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}

	var dummy = &ListNode{Next: head}
	var prev *ListNode
	current := dummy

	for left > 0 {
		prev = current
		current = current.Next
		left--
		right--
	}

	start := current
	then := start.Next

	for right > 0 {
		start.Next = then.Next
		then.Next = prev.Next
		prev.Next = then
		then = start.Next
		right--
	}

	return dummy.Next
}
