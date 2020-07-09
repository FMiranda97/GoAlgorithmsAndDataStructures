package main

import (
	"./linkedList"
	"fmt"
	"log"
)

type pessoa struct {
	id   int
	name string
}

/*func (p pessoa) getKey() string { //object oriented function
	return p.name
}*/

func (p pessoa) getKey() int { //object oriented function
	return p.id
}

func main() {
	list := linkedList.NewList()
	for i := 0; i < 5; i++ {
		var p pessoa
		fmt.Println("Inserir id:")
		_, _ = fmt.Scanf("%d", &p.id)
		fmt.Println("Inserir nome:")
		_, _ = fmt.Scanf("%s", &p.name)
		if err := linkedList.InsertList(&list, p); err != nil {
			log.Println(err)
		} else {
			linkedList.PrintList(list)
		}
	}
}
