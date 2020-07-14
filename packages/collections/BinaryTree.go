package collections

import (
	"errors"
)

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
func (tree *BinaryTree) Insert(cargo Sortable) error {
	var err error
	tree.root, err = insertBinaryTreeUtil(tree.root, cargo)
	if err == nil {
		tree.count++
	}
	return err
}

// Utility function for simple binary tree insertion
func insertBinaryTreeUtil(tree *treeNode, cargo Sortable) (*treeNode, error) {
	var err error
	if tree == nil { //reached insertion place
		newNode := treeNode{
			cargo: cargo,
		}
		return &newNode, nil
	}
	if cargo.CompareTo(tree.cargo) < 0 {
		tree.left, err = insertBinaryTreeUtil(tree.left, cargo)
	} else if cargo.CompareTo(tree.cargo) > 0 {
		tree.right, err = insertBinaryTreeUtil(tree.right, cargo)
	} else {
		return tree, errors.New("element with this key already exists")
	}
	return tree, err
}

// Method to remove cargo from simple binary tree
func (tree *BinaryTree) Remove(cargo Sortable) error {
	var err error
	tree.root, err = removeBinaryTreeUtil(tree.root, cargo)
	if err == nil {
		tree.count--
	}
	return err
}

// Utility method for cargo for cargo removal on simple binary tree
func removeBinaryTreeUtil(tree *treeNode, cargo Sortable) (*treeNode, error) {
	if tree == nil {
		return tree, errors.New("no element found with given key")
	}
	var err error
	if cargo.CompareTo(tree.cargo) < 0 {
		tree.left, err = removeBinaryTreeUtil(tree.left, cargo)
	} else if cargo.CompareTo(tree.cargo) > 0 {
		tree.right, err = removeBinaryTreeUtil(tree.right, cargo)
	} else {
		if tree.left == nil {
			return tree.right, nil
		} else if tree.right == nil {
			return tree.left, nil
		}
		var rightmost *treeNode
		for rightmost = tree.left; rightmost.right != nil; rightmost = rightmost.right {
		}
		tree.cargo = rightmost.cargo
		tree.left, err = removeBinaryTreeUtil(tree.left, rightmost.cargo)
	}
	return tree, err
}

// Method to retrieve cargo with a given key
func (tree BinaryTree) Get(cargo Sortable) (Sortable, error) {
	found, err := getBinaryTreeUtil(tree.root, cargo)
	if err == nil {
		return found.cargo, err
	}
	return nil, err
}

// Utility function for simple binary tree cargo retrieving
func getBinaryTreeUtil(tree *treeNode, cargo Sortable) (*treeNode, error) {
	if tree == nil {
		return nil, errors.New("no element found with given key")
	}
	if cargo.CompareTo(tree.cargo) < 0 {
		return getBinaryTreeUtil(tree.left, cargo)
	} else if cargo.CompareTo(tree.cargo) > 0 {
		return getBinaryTreeUtil(tree.right, cargo)
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

// Method to print simple binary tree layout.
// Passed argument defines how many " " characters there are in between tree levels
func (tree BinaryTree) PrintTree2D(spacing int) {
	// Pass initial space count as 0
	printTree2DUtil(tree.root, 0, spacing)
}
