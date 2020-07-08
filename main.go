package main

import (
	"fmt"
)

type pessoa struct {
	id   int
	name string
}

func comparePessoa(a pessoa, b pessoa) int {
	if a.id < b.id {
		return -1
	} else if a.id > b.id {
		return 1
	} else {
		return 0
	}
}

func main() {
	var list linkedList
	for i := 0; i < 5; i++ {
		var p pessoa
		fmt.Println("Inserir id:")
		fmt.Scanf("%d", &p.id)
		fmt.Println("Inserir nome:")
		fmt.Scanf("%s", &p.name)
		insertList(&list, p, comparePessoa)
		fmt.Println("----------------------")
		printList(list)
		fmt.Println("----------------------")
	}
}
