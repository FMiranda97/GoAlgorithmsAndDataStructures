package collections

import (
	"errors"
	"fmt"
)

// linked list object
type LinkedList struct {
	count int
	first *node
	last  *node
}

// linked list constructor
func NewLinkedList() LinkedList {
	return LinkedList{
		count: 0,
		first: nil,
		last:  nil,
	}
}

// method to insert cargo at end of linked list
func (list *LinkedList) Append(cargo interface{}) {
	newNode := node{
		cargo: cargo,
		next:  nil,
	}
	if list.last != nil {
		list.last.next = &newNode
		list.last = &newNode
	} else {
		list.first = &newNode
		list.last = &newNode
	}
	list.count++
}

// method to insert cargo at beginning of linked list
func (list *LinkedList) Prepend(cargo interface{}) {
	newNode := node{
		cargo: cargo,
		next:  nil,
	}
	newNode.next = list.first
	list.first = &newNode
	if newNode.next == nil {
		list.last = &newNode
	}
	list.count++
}

// method to insert cargo into linked list at a given position
func (list *LinkedList) Insert(cargo interface{}, index int) error {
	if list.count < index {
		return errors.New("index out of bounds")
	}
	newNode := node{
		cargo: cargo,
		next:  nil,
	}
	if index == 0 {
		newNode.next = list.first
		list.first = &newNode
	} else {
		aux := list.first
		for i := 1; i < index; i++ {
			aux = aux.next
		}
		newNode.next = aux.next
		aux.next = &newNode
	}
	if newNode.next == nil {
		list.last = &newNode
	}
	list.count++
	return nil
}

// removes cargo from linked list at index
func (list *LinkedList) Remove(index int) error {
	if list.count <= index {
		return errors.New("index out of bounds")
	}
	if index == 0 {
		list.first = list.first.next
		if list.first == nil {
			list.last = nil
		}
	} else {
		aux := list.first
		for i := 1; i < index; i++ {
			aux = aux.next
		}
		aux.next = aux.next.next
		if aux.next == nil {
			list.last = nil
		}
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

// method to retrieve cargo at a given position
// should not be used to iterate
func (list *LinkedList) Get(index int) (interface{}, error) {
	if list.count <= index {
		return nil, errors.New("index out of bounds")
	}
	aux := list.first
	for i := 0; i < index; i++ {
		aux = aux.next
	}
	return aux.cargo, nil
}

// method to return count of elements in linked list
func (list *LinkedList) Len() int {
	return list.count
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
