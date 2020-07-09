package linkedList

import (
	"errors"
	"fmt"
	"reflect"
)

func comparator(this interface{}, target interface{}) (int8, error) {
	if reflect.TypeOf(this) != reflect.TypeOf(target) {
		return 0, errors.New("compared parameters implement different interfaces")
	}
	switch this.(type) {
	case cargoLinkedListString:
		a := this.(cargoLinkedListString)
		b := target.(cargoLinkedListString)
		if a.GetKey() < b.GetKey() {
			return -1, nil
		} else if a.GetKey() > b.GetKey() {
			return 1, nil
		} else {
			return 0, nil
		}
	case cargoLinkedListInt:
		a := this.(cargoLinkedListInt)
		b := target.(cargoLinkedListInt)
		if a.GetKey() < b.GetKey() {
			return -1, nil
		} else if a.GetKey() > b.GetKey() {
			return 1, nil
		} else {
			return 0, nil
		}
	case cargoLinkedListInt64:
		a := this.(cargoLinkedListInt)
		b := target.(cargoLinkedListInt)
		if a.GetKey() < b.GetKey() {
			return -1, nil
		} else if a.GetKey() > b.GetKey() {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, errors.New("compared parameters implement unsupported interface")
	}
}

func InsertList(list *LinkedList, cargo interface{}) error {
	// check viability
	/*if *list != nil && reflect.TypeOf((*list).cargo) != reflect.TypeOf(cargo) {
		return errors.New("cargo is not of same type as contents previously inserted")
	}*/
	if !implementsLinkedListCargo(cargo) {
		return errors.New("cargo does not implement a supported interface")
	}
	var newNode node
	newNode.cargo = cargo
	//if insert at beginning
	if *list == nil {
		*list = &newNode
		return nil
	} else if comp, err := comparator((*list).cargo, newNode.cargo); comp > 0 {
		if err != nil {
			newNode.next = *list
			*list = &newNode
			return nil
		} else {
			return err
		}
	}
	//insert in order
	var i LinkedList
	for i = *list; i.next != nil; i = i.next {
		comp, err := comparator(i.next.cargo, newNode.cargo)
		if err != nil {
			return err
		}
		if comp > 0 {
			break
		}
	}
	newNode.next = i.next
	i.next = &newNode
	return nil
}

func PrintList(list LinkedList) {
	fmt.Println("----------------------")
	for i := list; i != nil; i = i.next {
		fmt.Println(*i)
	}
	fmt.Println("----------------------")
}

func NewList() LinkedList {
	return nil
}
