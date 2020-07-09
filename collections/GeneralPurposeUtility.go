package collections

import (
	"errors"
	"fmt"
	"reflect"
)

// general purpose structure for linear collections
type node struct {
	cargo interface{}
	next  *node
}

// general purpose structure for binary tree collections
type treeNode struct {
	cargo interface{}
	left  *treeNode
	right *treeNode
}

// compares keys of 2 cargos, return -1 if "this" less than "target", 0 if equal, 1 otherwise
func comparator(this interface{}, target interface{}) (int8, error) {
	if reflect.TypeOf(this) != reflect.TypeOf(target) {
		return 0, errors.New("compared parameters implement different interfaces")
	}
	switch this.(type) {
	case cargoString:
		a := this.(cargoString)
		b := target.(cargoString)
		if a.GetKey() < b.GetKey() {
			return -1, nil
		} else if a.GetKey() > b.GetKey() {
			return 1, nil
		} else {
			return 0, nil
		}
	case cargoInt:
		a := this.(cargoInt)
		b := target.(cargoInt)
		if a.GetKey() < b.GetKey() {
			return -1, nil
		} else if a.GetKey() > b.GetKey() {
			return 1, nil
		} else {
			return 0, nil
		}
	case cargoInt64:
		a := this.(cargoInt)
		b := target.(cargoInt)
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

// checks if passed cargo implements any of the required interfaces
func implementsLinkedListCargo(cargo interface{}) bool {
	if _, ok := cargo.(cargoString); ok {
		return true
	}
	if _, ok := cargo.(cargoInt); ok {
		return true
	}
	if _, ok := cargo.(cargoInt64); ok {
		return true
	}
	return false
}

// general purpose binary tree printing utility function
func printTreeUtil(tree *treeNode) {
	if tree == nil {
		return
	}
	printTreeUtil(tree.left)
	fmt.Println(tree.cargo)
	printTreeUtil(tree.right)
}

// general purpose 2D printing utility function
func printTree2DUtil(tree *treeNode, space int) {

	// Base case
	if tree == nil {
		return
	}

	// Increase distance between levels
	count := 8
	space += count

	// Process right child first
	printTree2DUtil(tree.right, space)

	// Print current node after space
	// count
	for i := count; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(tree.cargo)
	// Process left child
	printTree2DUtil(tree.left, space)
}

// Wrapper over printTree2DUtil()
func PrintTree2D(root *treeNode) {
	// Pass initial space count as 0
	printTree2DUtil(root, 0)
}
