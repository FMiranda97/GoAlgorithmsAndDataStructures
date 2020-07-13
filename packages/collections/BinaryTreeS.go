package collections

import (
	"errors"
)

// Simple binary tree object
type BinaryTreeS struct {
	root  *treeNode
	count int
}

// Simple binary tree constructor
func NewBinaryTreeS() BinaryTreeS {
	return BinaryTreeS{
		root:  nil,
		count: 0,
	}
}

// Method to insert cargo into binary tree
func (tree *BinaryTreeS) Insert(key string, cargo interface{}) error {
	var err error
	tree.root, err = insertBinaryTreeSUtil(tree.root, key, cargo)
	if err == nil {
		tree.count++
	}
	return err
}

// Utility function for simple binary tree insertion
func insertBinaryTreeSUtil(tree *treeNode, key string, cargo interface{}) (*treeNode, error) {
	var err error
	if tree == nil { //reached insertion place
		newNode := treeNode{
			key:   key,
			cargo: cargo,
		}
		return &newNode, nil
	}
	if key < tree.key {
		tree.left, err = insertBinaryTreeSUtil(tree.left, key, cargo)
	} else if key > tree.key {
		tree.right, err = insertBinaryTreeSUtil(tree.right, key, cargo)
	} else {
		return tree, errors.New("element with this key already exists")
	}
	return tree, err
}

// Method to remove cargo from simple binary tree
func (tree *BinaryTreeS) Remove(key string) error {
	var err error
	tree.root, err = removeBinaryTreeSUtil(tree.root, key)
	if err == nil {
		tree.count--
	}
	return err
}

// Utility method for cargo for cargo removal on simple binary tree
func removeBinaryTreeSUtil(tree *treeNode, key string) (*treeNode, error) {
	if tree == nil {
		return tree, errors.New("no element found with given key")
	}
	var err error
	if key < tree.key {
		tree.left, err = removeBinaryTreeSUtil(tree.left, key)
	} else if key > tree.key {
		tree.right, err = removeBinaryTreeSUtil(tree.right, key)
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
		tree.left, err = removeBinaryTreeSUtil(tree.left, rightmost.key)
	}
	return tree, err
}

// Method to retrieve cargo with a given key
func (tree BinaryTreeS) Get(key string) (interface{}, error) {
	found, err := getBinaryTreeSUtil(tree.root, key)
	if err == nil {
		return found.cargo, err
	}
	return nil, err
}

// Utility function for simple binary tree cargo retrieving
func getBinaryTreeSUtil(tree *treeNode, key string) (*treeNode, error) {
	if tree == nil {
		return nil, errors.New("no element found with given key")
	}
	if key < tree.key {
		return getBinaryTreeSUtil(tree.left, key)
	} else if key > tree.key {
		return getBinaryTreeSUtil(tree.right, key)
	} else {
		return tree, nil
	}
}

// Method returning number of elements in tree
func (tree BinaryTreeS) Count() int {
	return tree.count
}

// Method to print simple binary tree contents
func (tree BinaryTreeS) PrintTree() {
	printTreeUtil(tree.root)
}

// Method to print simple binary tree layout.
// Passed argument defines how many " " characters there are in between tree levels
func (tree BinaryTreeS) PrintTree2D(spacing int) {
	// Pass initial space count as 0
	printTree2DUtil(tree.root, 0, spacing)
}
