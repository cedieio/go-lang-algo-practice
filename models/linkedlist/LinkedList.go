package linkedlist

// Node represents a node in a singly linkedlist
type Node struct {
	Val  int
	Next *Node
}

func NewNode(val int) *Node {
	return &Node{val, nil}
}

func NewNodeWithNext(val int, n *Node) *Node {
	return &Node{val, n}
}

func (e *Node) GetNext() *Node {
	return e.Next
}

func (e *Node) Find(val int) *Node {
	curr := e

	for curr != nil {
		if curr.Val == val {
			return curr
		}

		curr = curr.Next
	}
	return nil
}

func (e *Node) Append(n *Node) *Node {
	if e == nil {
		return nil
	}

	curr := e
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = n
	return curr
}
