package stack

type Stack struct {
	head *Node
	Length int
}

type Node struct {
	value float64
	next *Node
}

func (s *Stack) Push(data float64) {
	newNode := &Node{data, nil}

	if s.head != nil {
		newNode.next = s.head
	}

	s.head = newNode
	s.Length++
}

func (s *Stack) Pop() float64 {
	if s.Length <= 0 {
		panic("stack is empty")
	}

	popped := *s.head
	s.head = popped.next

	s.Length--

	return popped.value
}
