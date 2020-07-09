package main

import (
	. "LinkedList/collections"
	"fmt"
	"log"
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
	/*file, err := os.Open("LinkedList/Inputs/NameId_10.txt")
	if err != nil {
		log.Println("File reading error", err)
		return
	}
	defer file.Close()
	linha := bufio.NewScanner(file)
	for linha.Scan() {

	}
	*/

	list := NewBinaryTree()
	var opcao int
	for i := 0; i < 10; i++ {
		var p pessoa
		fmt.Println("1 - Inserir")
		fmt.Println("2 - Remover")
		_, _ = fmt.Scanf("%d", &opcao)
		switch opcao {
		case 1:
			fmt.Println("Inserir id:")
			_, _ = fmt.Scanf("%d", &p.id)
			fmt.Println("Inserir nome:")
			_, _ = fmt.Scanf("%s", &p.name)
			if err := list.Insert(p); err != nil {
				log.Println(err)
			} else {
				list.PrintTree()
				fmt.Println("-----------")
				list.PrintTree2D()
				fmt.Println("-----------")
			}
			break
		case 2:
			/*info, err := linkedList.Pop(&list)
			if  err != nil {
				log.Println(err)
			} else {
				fmt.Println(info)
				linkedList.PrintList(list)
			}*/
			break
		default:
			break
		}
	}
}
