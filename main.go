package main

//todo check if structures use pointers
import (
	. "./packages/collections"
	"fmt"
	"math/rand"
)

type pessoa struct {
	id   int
	nome string
}

func (p pessoa) CompareTo(t interface{}) int8 { // necessary for collections and sorts usage
	target := t.(*pessoa)
	if p.id < target.id {
		return -1
	} else if p.id > target.id {
		return 1
	} else {
		return 0
	}
}

func createPessoa(id int, nome string) *pessoa { // utility
	return &pessoa{
		id:   id,
		nome: nome,
	}
}

func retrievePessoa(id int, tree AVLTree) *pessoa { // utility
	a, _ := tree.Get(pessoa{id: id})
	return a.(*pessoa)
}

func testCollection() {
	const size = 20
	tree := NewAVLTree()
	for i := 0; i < size; i++ {
		var p Sortable
		p = createPessoa(rand.Int()%(size*100), string(rand.Int()%100))
		_ = tree.Insert(p)
	}
	tree.PrintTree()
	p := retrievePessoa(1410, tree)
	p.nome = "isto é um nome dado"
	p = retrievePessoa(1410, tree)
	fmt.Println("------------------")
	fmt.Println(*p)
	fmt.Println(p.nome, p.id)
	fmt.Println("------------------")
	tree.PrintTree()
}

func main() {
	testCollection()
}
