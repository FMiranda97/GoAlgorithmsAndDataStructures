//package dataStructures provides constructs for storing data
package collections

import (
	"errors"
	"fmt"
)

// stack object
type Stack struct {
	count int
	first *node
}

// stack constructor
func NewStack() Stack {
	return Stack{
		count: 0,
		first: nil,
	}
}

// method to push cargo into stack
func (stack *Stack) Push(cargo interface{}) error {
	var err error
	stack.first, err = pushUtil(stack.first, cargo)
	return err
}

// utility function for stack pushing
func pushUtil(stack *node, cargo interface{}) (*node, error) {
	var newNode node
	newNode.cargo = cargo
	if stack == nil {
		return &newNode, nil
	} else {
		newNode.next = stack
		return &newNode, nil
	}

}

// method to pop cargo from stack
func (stack *Stack) Pop() (interface{}, error) {
	var cargo interface{}
	var err error
	cargo, stack.first, err = popUtil(stack.first)
	return cargo, err
}

// utility function for stack popping
func popUtil(stack *node) (interface{}, *node, error) {
	if stack == nil {
		return nil, nil, errors.New("empty stack")
	}
	//Stack has only 1 element
	return stack.cargo, stack.next, nil
}

//method to display stack contents
func (stack Stack) PrintStack() {
	aux := stack.first
	if aux == nil {
		fmt.Println("empty list")
	}
	for i := aux; i != nil; i = i.next {
		fmt.Println(*i)
	}
}
