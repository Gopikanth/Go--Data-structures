package main

import "fmt"

type Stack struct {
	items []int
}

//push

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

//pop

func (s *Stack) Pop() int {
	toRemove := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return toRemove
}

func main() {
	mystack := Stack{}
	fmt.Println(mystack)
	mystack.Push(100)
	mystack.Push(200)
	mystack.Push(300)
	mystack.Push(400)
	fmt.Println(mystack)
	mystack.Pop()
	fmt.Println(mystack)
}
