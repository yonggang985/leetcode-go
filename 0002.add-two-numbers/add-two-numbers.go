package Problem0002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	temp := result
	v, n := 0, 0
	for {
		v, n = add(l1, l2, n)
		temp.Val = v
		l1 = next(l1)
		l2 = next(l2)
		if l1 == nil && l2 == nil {
			break
		}
		temp.Next = &ListNode{}
		temp = temp.Next
	}
	if n == 1 {
		temp.Next = &ListNode{Val: n}
	}
	return result

}

func next(l *ListNode) *ListNode {
	if l != nil {
		return l.Next
	}
	return nil
}

func add(l1 *ListNode, l2 *ListNode, i int) (v, n int) {
	if l1 != nil {
		v += l1.Val
	}
	if l2 != nil {
		v += l2.Val
	}
	v += i
	if v > 9 {
		v -= 10
		n = 1
	}
	return
}
