package collections

import (
	"errors"
)

//todo verify documentation

// Simple binary tree object
type BinaryTree struct {
	root  *treeNode
	count int
}

// Simple binary tree constructor
func NewBinaryTree() BinaryTree {
	return BinaryTree{
		root:  nil,
		count: 0,
	}
}

// Method to insert cargo into binary tree
func (tree *BinaryTree) Insert(key string, cargo interface{}) error {
	var err error
	tree.root, err = insertBinaryTreeUtil(tree.root, key, cargo)
	if err == nil {
		tree.count++
	}
	return err
}

// Utility function for simple binary tree insertion
func insertBinaryTreeUtil(tree *treeNode, key string, cargo interface{}) (*treeNode, error) {
	var err error
	if tree == nil { //reached insertion place
		newNode := treeNode{
			key:   key,
			cargo: cargo,
		}
		return &newNode, nil
	}
	if key < tree.key {
		tree.left, err = insertBinaryTreeUtil(tree.left, key, cargo)
	} else if key > tree.key {
		tree.right, err = insertBinaryTreeUtil(tree.right, key, cargo)
	} else {
		return tree, errors.New("element with this key already exists")
	}
	return tree, err
}

// Method to remove cargo from simple binary tree
func (tree *BinaryTree) Remove(key string) error {
	var err error
	tree.root, err = removeBinaryTreeUtil(tree.root, key)
	if err == nil {
		tree.count--
	}
	return err
}

// Remove cargo from tree with a given key
func removeBinaryTreeUtil(tree *treeNode, key string) (*treeNode, error) {
	if tree == nil {
		return tree, errors.New("no element found with given key")
	}
	var err error
	if key < tree.key {
		tree.left, err = removeBinaryTreeUtil(tree.left, key)
	} else if key > tree.key {
		tree.right, err = removeBinaryTreeUtil(tree.right, key)
	} else {
		if tree.left == nil {
			return tree.right, nil
		} else if tree.right == nil {
			return tree.left, nil
		}
		var rightmost *treeNode
		for rightmost = tree.left; rightmost.right != nil; rightmost = rightmost.right {
		}
		tree.key = rightmost.key
		tree.cargo = rightmost.cargo
		tree.left, err = removeBinaryTreeUtil(tree.left, rightmost.key)
	}
	return tree, err
}

// Method to retrieve cargo with a given key
func (tree BinaryTree) Get(key string) (interface{}, error) {
	found, err := getBinaryTreeUtil(tree.root, key)
	if err == nil {
		return found.cargo, err
	}
	return nil, err
}

// Utility function for simple binary tree cargo retrieving
func getBinaryTreeUtil(tree *treeNode, key string) (*treeNode, error) {
	if tree == nil {
		return nil, errors.New("no element found with given key")
	}
	if key < tree.key {
		return getBinaryTreeUtil(tree.left, key)
	} else if key > tree.key {
		return getBinaryTreeUtil(tree.right, key)
	} else {
		return tree, nil
	}
}

// Method returning number of elements in tree
func (tree BinaryTree) Count() int {
	return tree.count
}

// Method to print simple binary tree contents
func (tree BinaryTree) PrintTree() {
	printTreeUtil(tree.root)
}

// Method to print simple binary tree layout
// Passed argument defines how much spacing there is between tree levels
func (tree BinaryTree) PrintTree2D(spacing int) {
	// Pass initial space count as 0
	printTree2DUtil(tree.root, 0, spacing)
}
