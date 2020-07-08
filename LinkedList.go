package main

import (
	"fmt"
)

type linkedList = *node

type node struct {
	cargo interface{}
	next *node
}

func insertList(list *linkedList, cargo interface{}, comparator func interface{}) {
	// TODO verify if arguments of same type
	var newNode node
	newNode.cargo = cargo
	if *list == nil {
		*list = &newNode
		return
	} else if comparator((*list).cargo, newNode.cargo) > 0 {
		newNode.next = *list
		*list = &newNode
		return
	}
	var i linkedList
	for i = *list; i.next != nil && comparator(i.next.cargo, newNode.cargo) < 0; i = i.next {
	}
	newNode.next = i.next
	i.next = &newNode
}

func printList(list linkedList) {
	for i := list; i != nil; i = i.next {
		fmt.Println(*i)
	}
}
