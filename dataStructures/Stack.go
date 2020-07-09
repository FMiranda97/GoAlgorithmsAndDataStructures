package dataStructures

import (
	"errors"
	"fmt"
)

type Stack struct {
	count int32
	first *node
}

func NewStack() Stack {
	return Stack{
		count: 0,
		first: nil,
	}
}

func (stack *Stack) Push(cargo interface{}) error {
	var err error
	stack.first, err = pushUtil(stack.first, cargo)
	return err
}

func pushUtil(stack *node, cargo interface{}) (*node, error) {
	if !implementsLinkedListCargo(cargo) {
		return stack, errors.New("cargo does not implement a supported interface")
	}
	var newNode node
	newNode.cargo = cargo
	if stack == nil {
		return &newNode, nil
	} else if _, err := comparator((*stack).cargo, newNode.cargo); err == nil {
		newNode.next = stack
		return &newNode, nil
	} else {
		return stack, err
	}
}

func (stack *Stack) Pop() (interface{}, error) {
	var cargo interface{}
	var err error
	cargo, stack.first, err = popUtil(stack.first)
	return cargo, err
}

func popUtil(stack *node) (interface{}, *node, error) {
	if stack == nil {
		return nil, nil, errors.New("empty stack")
	}
	//Stack has only 1 element
	return stack.cargo, stack.next, nil
}

func (stack Stack) PrintStack() {
	aux := stack.first
	fmt.Println("----------------------")
	if aux == nil {
		fmt.Println("empty list")
	}
	for i := aux; i != nil; i = i.next {
		fmt.Println(*i)
	}
	fmt.Println("----------------------")
}
