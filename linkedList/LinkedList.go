package linkedList

import (
	"errors"
	"fmt"
	"reflect"
)

func comparator(this interface{}, target interface{}) (int8, error) {
	if reflect.TypeOf(this) != reflect.TypeOf(target) {
		return 0, errors.New("compared parameters implement different interfaces")
	}
	switch this.(type) {
	case cargoLinkedListString:
		a := this.(cargoLinkedListString)
		b := target.(cargoLinkedListString)
		if a.GetKey() < b.GetKey() {
			return -1, nil
		} else if a.GetKey() > b.GetKey() {
			return 1, nil
		} else {
			return 0, nil
		}
	case cargoLinkedListInt:
		a := this.(cargoLinkedListInt)
		b := target.(cargoLinkedListInt)
		if a.GetKey() < b.GetKey() {
			return -1, nil
		} else if a.GetKey() > b.GetKey() {
			return 1, nil
		} else {
			return 0, nil
		}
	case cargoLinkedListInt64:
		a := this.(cargoLinkedListInt)
		b := target.(cargoLinkedListInt)
		if a.GetKey() < b.GetKey() {
			return -1, nil
		} else if a.GetKey() > b.GetKey() {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, errors.New("compared parameters implement unsupported interface")
	}
}

func InsertList(list *LinkedList, cargo interface{}) error {
	// check viability
	if !implementsLinkedListCargo(cargo) {
		return errors.New("cargo does not implement a supported interface")
	}
	var newNode node
	newNode.cargo = cargo
	//if insert at beginning
	if *list == nil {
		*list = &newNode
		return nil
	} else if comp, err := comparator((*list).cargo, newNode.cargo); comp > 0 || err != nil {
		if err == nil {
			newNode.next = *list
			*list = &newNode
			return nil
		} else {
			return err
		}
	}
	//insert in order
	var i LinkedList
	for i = *list; i.next != nil; i = i.next {
		comp, err := comparator(i.next.cargo, newNode.cargo)
		if err != nil {
			return err
		}
		if comp > 0 {
			break
		}
	}
	newNode.next = i.next
	i.next = &newNode
	return nil
}

func PrintList(list LinkedList) {
	fmt.Println("----------------------")
	if list == nil {
		fmt.Println("empty list")
	}
	for i := list; i != nil; i = i.next {
		fmt.Println(*i)
	}
	fmt.Println("----------------------")
}

func NewList() LinkedList {
	return nil
}

func NewTree() BinaryTree {
	return nil
}

func Push(stack *LinkedList, cargo interface{}) error {
	if !implementsLinkedListCargo(cargo) {
		return errors.New("cargo does not implement a supported interface")
	}
	var newNode node
	newNode.cargo = cargo
	if *stack == nil {
		*stack = &newNode
		return nil
	} else if comp, err := comparator((*stack).cargo, newNode.cargo); comp != 0 {
		if err == nil {
			newNode.next = *stack
			*stack = &newNode
			return nil
		} else {
			return err
		}
	}
	return nil
}

func Pop(stack *LinkedList) (interface{}, error) {
	if *stack == nil {
		return nil, errors.New("empty stack")
	}
	//Stack has only 1 element
	cargo := (*stack).cargo
	if (*stack).next == nil {
		*stack = nil
	} else {
		*stack = (*stack).next
	}
	return cargo, nil
}

//todo check errors
func InsertTree(tree BinaryTree, cargo interface{}) BinaryTree {
	var newNode treeNode
	newNode.cargo = cargo
	//if insert at beginning
	if tree == nil {
		return &newNode
	} else { // todo allow same key
		comp, _ := comparator((*tree).cargo, newNode.cargo)
		if comp < 0 {
			tree.right = InsertTree(tree.right, cargo)
		} else if comp > 0 {
			tree.left = InsertTree(tree.left, cargo)
		}
	}
	return tree
}

func PrintTree(tree BinaryTree) {
	if tree == nil {
		return
	}
	PrintTree(tree.left)
	fmt.Println(tree.cargo)
	PrintTree(tree.right)
}

func print2DUtil(tree BinaryTree, space int) {

	// Base case
	if tree == nil {
		return
	}

	// Increase distance between levels
	count := 8
	space += count

	// Process right child first
	print2DUtil(tree.right, space)

	// Print current node after space
	// count
	for i := count; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(tree.cargo)
	// Process left child
	print2DUtil(tree.left, space)
}

// Wrapper over print2DUtil()
func Print2D(tree BinaryTree) {
	// Pass initial space count as 0
	print2DUtil(tree, 0)
}
