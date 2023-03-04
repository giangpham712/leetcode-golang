package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	seen := make(map[int]bool)

	if head == nil {
		return nil
	}

	resultHead := ListNode{Val: head.Val}
	resultCurrent := &resultHead
	seen[head.Val] = true

	current := head
	for current.Next != nil {
		if seen[current.Next.Val] != true {
			seen[current.Next.Val] = true
			resultCurrent.Next = &ListNode{Val: current.Next.Val}
			resultCurrent = resultCurrent.Next
		}
		current = current.Next
	}

	return &resultHead
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	prevNode := head
	for node := head.Next; node != nil; node = node.Next {
		if node.Val == prevNode.Val {
			prevNode.Next = node.Next
		} else {
			prevNode = node
		}
	}

	return head
}
