package main

import (
	"fmt"
	"strconv"
)

const CACHE_SIZE = 5

type Node struct {
	Prev *Node
	Data string
	Next *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Hash map[string]*Node

type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Next = tail
	head.Prev = head

	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {

	if q.Length == 0 {
		fmt.Println("---------------- Cache Empty!! ---------------- ")
		return
	}

	itr := q.Head.Next
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", itr.Data)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		itr = itr.Next
	}
	fmt.Println("]")
}

func (c *Cache) Check(obj string) {
	node := &Node{}

	if val, ok := c.Hash[obj]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Data: obj}
	}

	c.Add(node)
	c.Hash[obj] = node

}

func (c *Cache) Add(n *Node) {
	fmt.Printf("Add: %v\n", n.Data)

	temp := c.Queue.Head.Next

	c.Queue.Head.Next = n
	n.Prev = c.Queue.Head
	n.Next = temp
	temp.Prev = n

	c.Queue.Length += 1
	if c.Queue.Length > CACHE_SIZE {
		c.Remove(c.Queue.Tail.Prev)
	}
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("Remove %v\n", n.Data)

	prev := n.Prev
	next := n.Next

	n.Prev.Next = next
	n.Next.Prev = prev

	c.Queue.Length -= 1
	delete(c.Hash, n.Data)

	return n
}

func main() {
	fmt.Println("Starting Go Cache")
	cache := NewCache()
	for _, word := range []string{"BMW", "Audi", "Land Rover", "General Motors", "Porsche", "Toyota", strconv.Itoa(12)} {
		// Check if elem already exists in Cache
		cache.Check(word)
		cache.Display()
	}

}
