package main

import "fmt"

type Queue struct {
	items []int
}

//Enque
func (q *Queue) Enque(i int) {
	q.items = append(q.items, i)

}

//Deque
func (q *Queue) Deque() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enque(100)
	myQueue.Enque(200)
	myQueue.Enque(300)
	fmt.Println(myQueue)
	myQueue.Deque()
	fmt.Println(myQueue)

}
