package collections

import (
	"fmt"
)

// general purpose structure for linear collections
type node struct {
	cargo interface{}
	next  *node
}

// general purpose structure for binary tree collections
type treeNode struct {
	key   string
	cargo interface{}
	left  *treeNode
	right *treeNode
}

// general purpose binary tree printing utility function
func printTreeUtil(tree *treeNode) {
	if tree == nil {
		return
	}
	printTreeUtil(tree.left)
	fmt.Println(tree.key, tree.cargo)
	printTreeUtil(tree.right)
}

// general purpose 2D printing utility function
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
	fmt.Println(tree.cargo)
	// Process left child
	printTree2DUtil(tree.left, space, spacing)
}
