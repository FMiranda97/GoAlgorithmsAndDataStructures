package collections

import (
	"fmt"
)

func init() {

}

// General purpose structure for linear collections
type node struct {
	cargo interface{}
	next  *node
}

// General purpose structure for binary tree collections with string as key
type treeNode struct {
	key    string
	cargo  interface{}
	height int
	isRed  bool
	left   *treeNode
	right  *treeNode
	parent *treeNode
}

// General purpose structure for binary tree collections with int as key
type treeNodeI struct {
	key    int64
	cargo  interface{}
	height int
	left   *treeNodeI
	right  *treeNodeI
	parent *treeNodeI
}

// General purpose binary tree printing utility function
func printTreeUtil(tree *treeNode) {
	if tree == nil {
		return
	}
	printTreeUtil(tree.left)
	fmt.Println(tree.key, tree.cargo)
	printTreeUtil(tree.right)
}

// General purpose binary tree printing utility function
func printTreeIUtil(tree *treeNodeI) {
	if tree == nil {
		return
	}
	printTreeIUtil(tree.left)
	fmt.Println(tree.key, tree.cargo)
	printTreeIUtil(tree.right)
}

// General purpose 2D printing utility function
func printTree2DUtil(tree *treeNode, space int, spacing int) {

	// Base case
	if tree == nil {
		return
	}

	// Increase distance between levels
	space += spacing

	// Process right child first
	printTree2DUtil(tree.right, space, spacing)

	// Print current node after space
	// count
	for i := spacing; i < space; i++ {
		fmt.Print(" ")
	}
	if tree.parent == nil {
		fmt.Println(tree.cargo, nil)
	} else {
		fmt.Println(tree.cargo, tree.parent.cargo)
	}
	// Process left child

	printTree2DUtil(tree.left, space, spacing)
}

// General purpose 2D printing utility function
func printTreeI2DUtil(tree *treeNodeI, space int, spacing int) {

	// Base case
	if tree == nil {
		return
	}

	// Increase distance between levels
	space += spacing

	// Process right child first
	printTreeI2DUtil(tree.right, space, spacing)

	// Print current node after space
	// count
	for i := spacing; i < space; i++ {
		fmt.Print(" ")
	}
	if tree.parent == nil {
		fmt.Println(tree.cargo, nil)
	} else {
		fmt.Println(tree.cargo, tree.parent.cargo)
	}
	// Process left child

	printTreeI2DUtil(tree.left, space, spacing)
}

// general purpose utility to find max of two int values
func getMax(value1 int, value2 int) int {
	if value1 >= value2 {
		return value1
	} else {
		return value2
	}
}
