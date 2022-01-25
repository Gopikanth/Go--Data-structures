package main

import "fmt"

const ArraySize = 7

type Hashtable struct {
	array [ArraySize]*bucket
}
type bucket struct {
	head *bucketnode
}
type bucketnode struct {
	key  string
	next *bucketnode
}

func main() {
	myhashtable := Init()
	list := []string{
		"Gopikanth",
		"Ravichandran",
		"shared",
	}
	for _, v := range list {
		myhashtable.Insert(v)
	}
	myhashtable.Delete("Bruce")

	myhashtable.Search("Gopikanth")

	myhashtable.Search("shared")
}

func (h *Hashtable) Insert(k string) {
	index := hash(k)
	h.array[index].insert(k)
}
func (h *Hashtable) Search(k string) bool {
	index := hash(k)
	return h.array[index].search(k)

}
func (h *Hashtable) Delete(k string) {
	index := hash(k)
	h.array[index].delete(k)
}

func Init() *Hashtable {
	result := &Hashtable{}
	for i := range result.array {
		result.array[i] = &bucket{}

	}
	return result
}
func hash(k string) int {
	sum := 0
	for _, v := range k {
		sum += int(v)
	}
	return sum % ArraySize
}
func (b *bucket) insert(k string) {

	if !b.search(k) {
		newNode := &bucketnode{key: k}
		newNode.next = b.head //setting up the head
		b.head = newNode
	} else {
		fmt.Println("Already exist")
	}
}
func (b *bucket) search(k string) bool {
	currentnode := b.head
	for currentnode != nil {
		if currentnode.key == k {
			return true
		}
		currentnode = currentnode.next
	}
	return false
}

func (b *bucket) delete(k string) { //Deleting the value in bucket
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previosnode := b.head
	for previosnode.next != nil {
		if previosnode.next.key == k {
			previosnode.next = previosnode.next.next
		}
		previosnode = previosnode.next
	}
}
