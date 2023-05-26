package main

import "fmt"

const (
	SIZE = 5
)

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Hash map[string]*Node

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")

}

type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewCache() Cache {
	return Cache{
		Queue: NewQueue(),
		Hash:  Hash{},
	}
}

func (c *Cache) Check(s string) {
	node := &Node{}
	if v, ok := c.Hash[s]; ok {
		node = c.Remove(v)
	} else {
		node = &Node{Val: s}
	}
	c.Add(node)
	c.Hash[s] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove: %s\n", n.Val)
	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left
	c.Queue.Length -= 1
	delete(c.Hash, n.Val)
	return n
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("add: %s\n", n.Val)
	t := c.Queue.Head.Right

	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = t
	t.Left = n

	c.Queue.Length++
	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func main() {
	fmt.Println("START CACHE")
	cache := NewCache()
	for _, v := range []string{"Emory", "Lin", "Gao", "Apple", "Go", "Java", "Rust", "C", "Emory"} {
		cache.Check(v)
		cache.Display()
	}
}
