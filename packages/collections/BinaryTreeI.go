package collections

import (
	"errors"
)

// Simple binary tree object
type BinaryTreeI struct {
	root  *treeNodeI
	count int
}

// Simple binary tree constructor
func NewBinaryTreeI() BinaryTreeI {
	return BinaryTreeI{
		root:  nil,
		count: 0,
	}
}

// Method to insert cargo into binary tree
func (tree *BinaryTreeI) Insert(key int64, cargo interface{}) error {
	var err error
	tree.root, err = insertBinaryTreeIUtil(tree.root, key, cargo)
	if err == nil {
		tree.count++
	}
	return err
}

// Utility function for simple binary tree insertion
func insertBinaryTreeIUtil(tree *treeNodeI, key int64, cargo interface{}) (*treeNodeI, error) {
	var err error
	if tree == nil { //reached insertion place
		newNode := treeNodeI{
			key:   key,
			cargo: cargo,
		}
		return &newNode, nil
	}
	if key < tree.key {
		tree.left, err = insertBinaryTreeIUtil(tree.left, key, cargo)
	} else if key > tree.key {
		tree.right, err = insertBinaryTreeIUtil(tree.right, key, cargo)
	} else {
		return tree, errors.New("element with this key already exists")
	}
	return tree, err
}

// Method to remove cargo from simple binary tree
func (tree *BinaryTreeI) Remove(key int64) error {
	var err error
	tree.root, err = removeBinaryTreeIUtil(tree.root, key)
	if err == nil {
		tree.count--
	}
	return err
}

// Utility method for cargo for cargo removal on simple binary tree
func removeBinaryTreeIUtil(tree *treeNodeI, key int64) (*treeNodeI, error) {
	if tree == nil {
		return tree, errors.New("no element found with given key")
	}
	var err error
	if key < tree.key {
		tree.left, err = removeBinaryTreeIUtil(tree.left, key)
	} else if key > tree.key {
		tree.right, err = removeBinaryTreeIUtil(tree.right, key)
	} else {
		if tree.left == nil {
			return tree.right, nil
		} else if tree.right == nil {
			return tree.left, nil
		}
		var rightmost *treeNodeI
		for rightmost = tree.left; rightmost.right != nil; rightmost = rightmost.right {
		}
		tree.key = rightmost.key
		tree.cargo = rightmost.cargo
		tree.left, err = removeBinaryTreeIUtil(tree.left, rightmost.key)
	}
	return tree, err
}

// Method to retrieve cargo with a given key
func (tree BinaryTreeI) Get(key int64) (interface{}, error) {
	found, err := getBinaryTreeIUtil(tree.root, key)
	if err == nil {
		return found.cargo, err
	}
	return nil, err
}

// Utility function for simple binary tree cargo retrieving
func getBinaryTreeIUtil(tree *treeNodeI, key int64) (*treeNodeI, error) {
	if tree == nil {
		return nil, errors.New("no element found with given key")
	}
	if key < tree.key {
		return getBinaryTreeIUtil(tree.left, key)
	} else if key > tree.key {
		return getBinaryTreeIUtil(tree.right, key)
	} else {
		return tree, nil
	}
}

// Method returning number of elements in tree
func (tree BinaryTreeI) Count() int {
	return tree.count
}

// Method to print simple binary tree contents
func (tree BinaryTreeI) PrintTree() {
	printTreeIUtil(tree.root)
}

// Method to print simple binary tree layout.
// Passed argument defines how many " " characters there are in between tree levels
func (tree BinaryTreeI) PrintTree2D(spacing int) {
	// Pass initial space count as 0
	printTreeI2DUtil(tree.root, 0, spacing)
}
