package collections

import (
	"errors"
)

//todo count elements

// simple binary tree object
type BinaryTree struct {
	root  *treeNode
	count int32
}

// simple binary tree constructor
func NewBinaryTree() BinaryTree {
	return BinaryTree{
		root:  nil,
		count: 0,
	}
}

// method to insert cargo into binary tree
func (tree *BinaryTree) Insert(cargo interface{}) error {
	var err error
	tree.root, err = insertBinaryTreeUtil(tree.root, cargo)
	if err == nil {
		tree.count++
	}
	return err
}

// utility function for simple binary tree insertion
func insertBinaryTreeUtil(tree *treeNode, cargo interface{}) (*treeNode, error) {
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
			tree.right, err = insertBinaryTreeUtil(tree.right, cargo)
		} else if comp > 0 {
			tree.left, err = insertBinaryTreeUtil(tree.left, cargo)
		} else {
			return tree, errors.New("element with this key already exists")
		}
	}
	return tree, nil
}

// method to print simple binary tree contents
func (tree BinaryTree) PrintTree() {
	printTreeUtil(tree.root)
}
