package collections

import (
	"errors"
	"fmt"
)

//todo get errors to print stack traces

// linked list object
type LinkedList struct {
	count int
	first *node
}

// linked list constructor
func NewLinkedList() LinkedList {
	return LinkedList{
		count: 0,
		first: nil,
	}
}

// method to insert cargo into linked list
func (list *LinkedList) Insert(cargo interface{}) error {
	var err error
	list.first, err = insertLinkedListUtil(list.first, cargo)
	if err == nil {
		list.count++
	}
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

// removes cargo from linked list at index
func (list *LinkedList) Remove(index int) error {
	if list.count <= index {
		return errors.New("index out of bounds")
	}
	if index == 0 {
		list.first = list.first.next
	} else {
		aux := list.first
		for i := 1; i < index; i++ {
			aux = aux.next
		}
		aux.next = aux.next.next
	}
	list.count--
	return nil
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

// returns an iterator function. Each call to the return function returns the next element
// once last element is reached returns nil forever
func (list *LinkedList) GetIterator() func() interface{} {
	elem := list.first
	return func() interface{} {
		if elem == nil {
			return nil
		}
		aux := elem.cargo
		elem = elem.next
		return aux
	}
}

// builds and returns an array with cargos present in LinkedList
func (list *LinkedList) GetArray() interface{} {
	arr := make([]interface{}, list.count)
	it := list.GetIterator()
	for elem, i := it(), 0; elem != nil; elem, i = it(), i+1 {
		arr[i] = elem
	}
	return arr
}
