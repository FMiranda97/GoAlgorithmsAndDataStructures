package main

import (
	"./dataStructures"
	"fmt"
	"log"
	"math/rand"
)

type pessoa struct {
	id   int
	name string
}

func (p pessoa) GetKey() string { //object oriented function
	return p.name
}

/*func (p pessoa) GetKey() int { //object oriented function
	return p.id
}*/

func main() {
	pilha := dataStructures.NewStack()
	nome := ""
	for i := 0; i < 5; i++ {
		nome = nome + "a"
		p := pessoa{
			id:   rand.Int() % 91,
			name: nome,
		}
		_ = pilha.Push(p)
	}
	pilha.PrintStack()
	for i := 0; i < 6; i++ {
		if c, err := pilha.Pop(); err == nil {
			fmt.Println(c)
		} else {
			log.Println(err)
		}
	}
}
