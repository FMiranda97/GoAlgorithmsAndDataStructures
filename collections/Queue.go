package collections

import (
	"errors"
	"fmt"
)

//todo verify documentation

// linked list object
type Queue struct {
	count int
	first *node
	last  *node
}

// linked list constructor
func NewQueue() Queue {
	return Queue{
		count: 0,
		first: nil,
		last:  nil,
	}
}

// method to insert cargo at end of linked list
func (queue *Queue) Append(cargo interface{}) {
	newNode := node{
		cargo: cargo,
		next:  nil,
	}
	if queue.last != nil {
		queue.last.next = &newNode
		queue.last = &newNode
	} else {
		queue.first = &newNode
		queue.last = &newNode
	}
	queue.count++
}

// method to pop cargo from queue
func (queue *Queue) Pop() (cargo interface{}, err error) {
	if queue.count == 0 {
		return nil, errors.New("empty queue")
	}
	cargo = queue.first.cargo
	queue.first = queue.first.next
	if queue.first == nil {
		queue.last = nil
	}
	queue.count--
	return cargo, nil
}

// method to display linked list contents
func (queue Queue) PrintQueue() {
	aux := queue.first
	if aux == nil {
		fmt.Println("empty queue")
	}
	for i := aux; i != nil; i = i.next {
		fmt.Println(*i)
	}
}

func (queue *Queue) Len() int {
	return queue.count
}
