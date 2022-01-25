package main

import (
	"fmt"
)

type Graph struct {
	vertices []*Vertex
}
type Vertex struct {
	key       int
	adajacent []*Vertex
}

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func (g *Graph) Addvertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added it is exist already", k)
		fmt.Println(err.Error())
	}
	g.vertices = append(g.vertices, &Vertex{key: k})
}

func (g *Graph) Addedge(From, To int) {
	fromvertex := g.getVertex(From)
	tovertex := g.getVertex(To)

	if fromvertex == nil || tovertex == nil {
		err := fmt.Errorf("invalid edge %v --->%v", From, To)
		fmt.Println(err.Error())
	} else if contains(fromvertex.adajacent, To) {
		err := fmt.Errorf("Existing edge %v --->%v", From, To)
		fmt.Println(err.Error())
	} else {
		fromvertex.adajacent = append(fromvertex.adajacent, tovertex)
	}

}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v :", v.key)
		for _, v := range v.adajacent {
			fmt.Printf(" %v", v.key)
		}
	}
	fmt.Println()
}
func (g Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func main() {
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.Addvertex(i)
	}

	test.Addedge(4, 1)
	test.Addedge(6, 3)
	test.Addvertex(0)
	test.Addedge(6, 1)
	test.Print()
}
