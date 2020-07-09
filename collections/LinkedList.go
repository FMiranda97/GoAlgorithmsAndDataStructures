package collections

import (
	"errors"
	"fmt"
)

//TODO count elements
//todo iterator
//todo convert to array
//todo construct from array

type node struct {
	cargo interface{}
	next  *node
}

type LinkedList struct {
	count int32
	first *node
}

func insertLinkedListUtil(list *node, cargo interface{}) (*node, error) {
	// check viability
	if !implementsLinkedListCargo(cargo) {
		return list, errors.New("cargo does not implement a supported interface")
	}
	var newNode node
	newNode.cargo = cargo
	//if insert at beginning
	if list == nil {
		return &newNode, nil
	} else if comp, err := comparator((*list).cargo, newNode.cargo); comp > 0 || err != nil {
		if err == nil {
			newNode.next = list
			return &newNode, nil
		} else {
			return list, err
		}
	}
	//insert in order
	var i *node
	for i = list; i.next != nil; i = i.next {
		comp, err := comparator(i.next.cargo, newNode.cargo)
		if err != nil {
			return list, err
		}
		if comp > 0 {
			break
		}
	}
	newNode.next = i.next
	i.next = &newNode
	return list, nil
}

func (list LinkedList) PrintList() {
	aux := list.first
	fmt.Println("----------------------")
	if aux == nil {
		fmt.Println("empty list")
	}
	for i := aux; i != nil; i = i.next {
		fmt.Println(*i)
	}
	fmt.Println("----------------------")
}

func NewList() LinkedList {
	return LinkedList{
		count: 0,
		first: nil,
	}
}

func (list *LinkedList) Insert(cargo interface{}) error {
	var err error
	list.first, err = insertLinkedListUtil(list.first, cargo)
	return err
}
