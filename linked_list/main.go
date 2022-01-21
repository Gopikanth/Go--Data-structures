package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	lenght int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.lenght++
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: 21}
	node2 := &node{data: 53}
	node3 := &node{data: 45}
	node4 := &node{data: 23}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	fmt.Println(mylist)
}
