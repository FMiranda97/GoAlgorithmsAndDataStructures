package collections

import (
	"errors"
	"fmt"
)

// Deque object
type Deque struct {
	count int
	first *node
	last  *node
}

// Deque constructor
func NewDeque() Deque {
	return Deque{
		count: 0,
		first: nil,
		last:  nil,
	}
}

// Method to insert cargo at end of deque
func (deque *Deque) Append(cargo *interface{}) {
	newNode := node{
		cargo: cargo,
		next:  nil,
	}
	if deque.last != nil {
		deque.last.next = &newNode
		deque.last = &newNode
	} else {
		deque.first = &newNode
		deque.last = &newNode
	}
	deque.count++
}

// Method to insert cargo at beginning of deque
func (deque *Deque) Prepend(cargo *interface{}) {
	newNode := node{
		cargo: cargo,
		next:  nil,
	}
	newNode.next = deque.first
	deque.first = &newNode
	if newNode.next == nil {
		deque.last = &newNode
	}
	deque.count++
}

// Method to pop cargo from deque's front
func (deque *Deque) Pop() (cargo interface{}, err error) {
	if deque.count == 0 {
		return nil, errors.New("empty deque")
	}
	cargo = deque.first.cargo
	deque.first = deque.first.next
	if deque.first == nil {
		deque.last = nil
	}
	deque.count--
	return cargo, nil
}

// Method to pop cargo from deque's back
func (deque *Deque) PopEnd() (cargo interface{}, err error) {
	if deque.count == 0 {
		return nil, errors.New("empty deque")
	}
	cargo = deque.last.cargo
	if deque.first == deque.last {
		deque.first = nil
		deque.last = nil
	} else {
		aux := deque.first
		for aux.next != deque.last {
			aux = aux.next
		}
		deque.last = aux
		aux.next = nil
	}
	deque.count--
	return cargo, nil
}

// Method to return count of elements in deque
func (deque Deque) Len() int {
	return deque.count
}

// Method to display deque contents
func (deque Deque) PrintDeque() {
	aux := deque.first
	if aux == nil {
		fmt.Println("empty deque")
	}
	for i := aux; i != nil; i = i.next {
		fmt.Println(*i)
	}
}
