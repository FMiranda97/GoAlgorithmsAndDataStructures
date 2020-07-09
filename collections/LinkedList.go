package collections

import (
	"errors"
	"fmt"
)

//TODO count elements
//todo iterator
//todo convert to array
//todo construct from array
//todo remove

// linked list object
type LinkedList struct { //todo check if should be lower case
	count int32
	first *node
}

// linked list constructor
func NewList() LinkedList {
	return LinkedList{
		count: 0,
		first: nil,
	}
}

// method to insert cargo into linked list
func (list *LinkedList) Insert(cargo interface{}) error {
	var err error
	list.first, err = insertLinkedListUtil(list.first, cargo)
	return err
}

// utility function for linked list insertion
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

// method to display linked list contents
func (list LinkedList) PrintList() {
	aux := list.first
	if aux == nil {
		fmt.Println("empty list")
	}
	for i := aux; i != nil; i = i.next {
		fmt.Println(*i)
	}
}
