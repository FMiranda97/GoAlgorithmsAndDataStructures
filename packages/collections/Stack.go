// Package collections provides constructs for storing and searching data
package collections

import (
	"errors"
	"fmt"
)

// Stack object
type Stack struct {
	count int
	first *node
}

// Stack constructor
func NewStack() Stack {
	return Stack{
		count: 0,
		first: nil,
	}
}

// Method to push cargo into stack
func (stack *Stack) Push(cargo *interface{}) {
	stack.first = pushUtil(stack.first, cargo)
	stack.count++
}

// Utility function for stack pushing
func pushUtil(stack *node, cargo *interface{}) *node {
	var newNode node
	newNode.cargo = cargo
	if stack == nil {
		return &newNode
	} else {
		newNode.next = stack
		return &newNode
	}

}

// Method to pop cargo from stack
func (stack *Stack) Pop() (interface{}, error) {
	var cargo interface{}
	var err error
	cargo, stack.first, err = popStackUtil(stack.first)
	if err == nil {
		stack.count--
	}
	return cargo, err
}

// Utility function for stack popping
func popStackUtil(stack *node) (interface{}, *node, error) {
	if stack == nil {
		return nil, nil, errors.New("empty stack")
	}
	//Stack has only 1 element
	return stack.cargo, stack.next, nil
}

// Method to return count of elements in stack
func (stack Stack) Len() int {
	return stack.count
}

// Method to display stack contents
func (stack Stack) PrintStack() {
	aux := stack.first
	if aux == nil {
		fmt.Println("empty list")
	}
	for i := aux; i != nil; i = i.next {
		fmt.Println(*i)
	}
}
