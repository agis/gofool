package stack

type Stack struct {
	head *Node
	Length int
}

type Node struct {
	value int
	next *Node
}

func (s *Stack) Push(data int) {
	newNode := &Node{data, nil}

	if s.head != nil {
		newNode.next = s.head
	}

	s.head = newNode
	s.Length++
}

func (s *Stack) Pop() int {
	if s.Length <= 0 {
		panic("stack is empty")
	}

	popped := *s.head
	s.head = popped.next

	s.Length--

	return popped.value
}
