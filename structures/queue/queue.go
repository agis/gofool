package queue

import "fmt"

type Node struct {
  data interface{}
  prev *Node
  next *Node
}

type Yo interface {
	A() int
}

type Queue struct {
  head *Node
  tail *Node
  length int
}

func (q *Queue) enqueue(data Node) *Node {
  node := &Node{data.data, nil, q.head}

  if q.head != nil {
    q.head.prev = node
  }

  q.head = node

  if q.tail == nil {
    q.tail = node
  }

  q.length++

  return node
}

func (q *Queue) dequeue() interface{} {
  if q.tail == nil {
    panic("tried to dequeue empty list!")
  }

  data := q.tail.data
  q.tail = q.tail.prev

  if q.tail == nil { // there was 1 node before, so queue is empty now
    q.head = nil
  } else {
    q.tail.next = nil
  }

  q.length--

  return data
}

func (q *Queue) print() {
  fmt.Printf("(%d) ", q.length)

  for p := q.head; p != nil; p = p.next {
    fmt.Printf("%d", p.data)

    if p.next != nil {
      fmt.Printf(" <-> ")
    }
  }

  fmt.Println()
}
