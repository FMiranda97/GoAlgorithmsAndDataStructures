package main

import (
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

//TODO make inserts avalaible outside package through upper case

func main() {
	var list linkedList
	for i := 0; i < 5; i++ {
		var p pessoa
		fmt.Println("Inserir id:")
		_, _ = fmt.Scanf("%d", &p.id)
		fmt.Println("Inserir nome:")
		_, _ = fmt.Scanf("%s", &p.name)
		if err := insertList(&list, p); err != nil {
			log.Println(err)
		} else {
			printList(list)
		}
	}
}
