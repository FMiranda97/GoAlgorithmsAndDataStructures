package collections

import (
	"errors"
	"fmt"
)

// Linked list object
type OrderedLinkedList struct {
	count int
	first *nodeOrder
	last  *nodeOrder
}

// Linked list constructor
func NewOrderedLinkedList() OrderedLinkedList {
	return OrderedLinkedList{
		count: 0,
		first: nil,
	}
}

// Method to insert cargo into linked list at a given position
func (list *OrderedLinkedList) Insert(cargo Sortable) error {
	newNode := nodeOrder{
		cargo: cargo,
		next:  nil,
	}
	if list.first == nil {
		list.count++
		list.first = &newNode
		return nil
	} else if cargo.CompareTo(list.first.cargo) <= 0 {
		if cargo.CompareTo(list.first.cargo) == 0 {
			return errors.New("element already exists")
		} else {
			newNode.next = list.first
			list.first = &newNode
			return nil
		}
	}
	var aux *nodeOrder
	for aux = list.first; aux.next != nil && cargo.CompareTo(aux.next.cargo) > 0; aux = aux.next {
	}
	if aux.next != nil && cargo.CompareTo(aux.next.cargo) == 0 {
		return errors.New("element already exists")
	}
	newNode.next = aux.next
	aux.next = &newNode
	list.count++
	return nil
}

// Removes cargo from linked list at index
func (list *OrderedLinkedList) Remove(cargo Sortable) error {
	var aux *nodeOrder
	if list.first == nil {
		return errors.New("empty list")
	}
	if cargo.CompareTo(list.first.cargo) == 0 {
		list.first = list.first.next
		list.count--
		return nil
	}
	for aux = list.first; aux.next != nil && cargo.CompareTo(aux.next.cargo) > 0; aux = aux.next {
	}
	if aux.next != nil && cargo.CompareTo(aux.next.cargo) == 0 {
		aux.next = aux.next.next
		list.count--
		return nil
	} else {
		return errors.New("cargo not found")
	}
}

// Method to display linked list contents
func (list OrderedLinkedList) PrintList() {
	aux := list.first
	if aux == nil {
		fmt.Println("empty list")
	}
	for i := aux; i != nil; i = i.next {
		fmt.Println(i.cargo)
	}
}

// Method to retrieve cargo at a given position.
// Should not be used to iterate
func (list *OrderedLinkedList) Get(cargo Sortable) (Sortable, error) {
	for aux := list.first; aux != nil; aux = aux.next {
		if cargo.CompareTo(aux.cargo) == 0 {
			return aux.cargo, nil
		}
	}
	return nil, errors.New("cargo not found")
}

// Method to return count of elements in linked list
func (list OrderedLinkedList) Len() int {
	return list.count
}

// Returns an iterator function. Each call to the return function returns the next element.
// Once last element is reached returns nil forever
func (list OrderedLinkedList) GetIterator() func() Sortable {
	elem := list.first
	return func() Sortable {
		if elem == nil {
			return nil
		}
		aux := elem.cargo
		elem = elem.next
		return aux
	}
}

// Builds and returns an array with cargos present in OrderedLinkedList
func (list OrderedLinkedList) GetArray() []Sortable {
	arr := make([]Sortable, list.count)
	it := list.GetIterator()
	for elem, i := it(), 0; elem != nil; elem, i = it(), i+1 {
		arr[i] = elem
	}
	return arr
}
