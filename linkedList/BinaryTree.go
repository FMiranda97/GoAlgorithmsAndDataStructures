package linkedList

import (
	"errors"
	"fmt"
)

func (tree BinaryTree) PrintTree() {
	printTreeUtil(tree.root)
}

func printTreeUtil(tree *treeNode) {
	if tree == nil {
		return
	}
	printTreeUtil(tree.left)
	fmt.Println(tree.cargo)
	printTreeUtil(tree.right)
}

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
func (tree BinaryTree) PrintTree2D() {
	// Pass initial space count as 0
	printTree2DUtil(tree.root, 0)
}

//todo count elements

type treeNode struct {
	cargo interface{}
	left  *treeNode
	right *treeNode
}

type BinaryTree struct {
	root  *treeNode
	count int32
}

func NewBinaryTree() BinaryTree {
	return BinaryTree{
		root:  nil,
		count: 0,
	}
}

func (tree *BinaryTree) Insert(cargo interface{}) error {
	var err error
	tree.root, err = insertUtil(tree.root, cargo)
	if err == nil {
		tree.count++
	}
	return err
}

func insertUtil(tree *treeNode, cargo interface{}) (*treeNode, error) {
	if !implementsLinkedListCargo(cargo) {
		return tree, errors.New("cargo does not implement a supported interface")
	}
	var newNode treeNode
	newNode.cargo = cargo
	//if insert at beginning
	if tree == nil {
		return &newNode, nil
	} else {
		comp, err := comparator((*tree).cargo, newNode.cargo)
		if err != nil {
			return tree, err
		}
		if comp < 0 {
			tree.right, err = insertUtil(tree.right, cargo)
		} else if comp > 0 {
			tree.left, err = insertUtil(tree.left, cargo)
		} else {
			return tree, errors.New("element with this key already exists")
		}
	}
	return tree, nil
}
