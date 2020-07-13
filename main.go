package main

import (
	. "./packages/collections"
	"fmt"
	"math/rand"
)

type pessoa struct {
	id int
}

func (p pessoa) CompareTo(t interface{}) int8 {
	target := t.(pessoa)
	if p.id < target.id {
		return -1
	} else if p.id > target.id {
		return 1
	} else {
		return 0
	}
}

func testCollection() {
	const size = 20
	tree := NewAVLTree()
	for i := 0; i < size; i++ {
		p := pessoa{
			id: rand.Int() % (size * 100),
		}
		_ = tree.Insert(p)
	}
	tree.PrintTree2D(10)
	fmt.Println("----------------------")
	got, _ := tree.Get(pessoa{id: 836})
	fmt.Println(got)
	for i := 0; i < size*100; i++ {
		_ = tree.Remove(pessoa{id: rand.Int() % (size * 100)})
	}
	fmt.Println("----------------------")
	tree.PrintTree2D(10)
}

func main() {
	testCollection()
}
