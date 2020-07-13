package collections

import (
	"errors"
	"fmt"
)

// Red-Black tree object
type RBTree struct {
	root  *treeNode
	count int
}

// Red-Black tree constructor
func NewRBTreeS() RBTree {
	return RBTree{
		root:  nil,
		count: 0,
	}
}

// Method to insert cargo into Red-Black tree
func (tree *RBTree) Insert(key string, cargo interface{}) (err error) {
	tree.root, err = insertRBTreeUtil(tree.root, key, cargo)
	if err == nil {
		tree.count++
	}
	return
}

// Utility function for Red-Black tree insertion
func actualRBInsertion(tree *treeNode, key string, newNode *treeNode) (*treeNode, error) {
	var err error
	if tree == nil {
		return newNode, nil
	}
	if key < tree.key {
		tree.left, err = actualRBInsertion(tree.left, key, newNode)
		tree.left.parent = tree
	} else if key > tree.key {
		tree.right, err = actualRBInsertion(tree.right, key, newNode)
		tree.right.parent = tree
	} else {
		return tree, errors.New("element with this key already exists")
	}
	return tree, err
}

// Utility function to perform Red-Black tree insertion
func insertRBTreeUtil(tree *treeNode, key string, cargo interface{}) (*treeNode, error) {
	var err error
	newNode := &treeNode{
		key:   key,
		cargo: cargo,
		isRed: true,
	}
	tree, err = actualRBInsertion(tree, key, newNode)
	/*auto balance tree*/
	var uncleNode *treeNode
	for newNode.parent != nil && newNode.isRed && newNode.parent.isRed {
		if newNode.parent == newNode.parent.parent.left {
			uncleNode = newNode.parent.parent.right
			if uncleNode != nil && uncleNode.isRed {
				newNode.parent.isRed = false
				uncleNode.isRed = false
				newNode.parent.parent.isRed = true
				newNode = newNode.parent.parent
			} else if uncleNode == nil || !uncleNode.isRed {
				rotRec(newNode)
			}
		} else {
			uncleNode = newNode.parent.parent.left
			if uncleNode != nil && uncleNode.isRed {
				newNode.parent.isRed = false
				uncleNode.isRed = false
				newNode.parent.parent.isRed = true
				newNode = newNode.parent.parent
			} else if uncleNode == nil || !uncleNode.isRed {
				rotRec(newNode)
			}
		}
	}

	for ; tree != nil && tree.parent != nil; tree = tree.parent {
	}
	if tree != nil {
		tree.isRed = false
	}
	return tree, err
}

func rotRec(treeNode *treeNode) *treeNode {
	aux := treeNode
	treeNode = treeNode.parent
	if treeNode.parent.left != nil && treeNode.right != nil && treeNode.right == aux && treeNode.parent.left == treeNode {
		leftRotateRB(treeNode)
		treeNode.parent.isRed = false
		treeNode.parent.parent.isRed = true
		rightRotateRB(treeNode.parent.parent)
	} else if treeNode.parent.right != nil && treeNode.left != nil && treeNode.left == aux && treeNode.parent.right == treeNode {
		rightRotateRB(treeNode)
		treeNode.parent.isRed = false
		treeNode.parent.parent.isRed = true
		leftRotateRB(treeNode.parent.parent)
	} else if treeNode.right != nil && treeNode.right == aux && treeNode.right.parent == treeNode && treeNode.parent.right == treeNode {
		leftRotateRB(treeNode.parent)
		treeNode.isRed = false
		treeNode.left.isRed = true
	} else if treeNode.left != nil && treeNode.left == aux && treeNode.left.parent == treeNode && treeNode.parent.left != nil {
		rightRotateRB(treeNode.parent)
		treeNode.isRed = false
		treeNode.right.isRed = true
	}
	return treeNode
}

func rightRotateRB(treeNode *treeNode) *treeNode {
	leftNode := treeNode.left
	leftRightNode := leftNode.right

	leftNode.right = treeNode
	treeNode.left = leftRightNode

	if treeNode.parent != nil {
		if treeNode.parent.right == treeNode {
			treeNode.parent.right = leftNode
		} else {
			treeNode.parent.left = leftNode
		}
	}

	leftNode.parent = treeNode.parent
	treeNode.parent = leftNode
	if leftRightNode != nil {
		leftRightNode.parent = treeNode
	}
	return leftNode
}

func leftRotateRB(treeNode *treeNode) *treeNode {
	rightNode := treeNode.right
	rightLeftNode := rightNode.left

	rightNode.left = treeNode
	treeNode.right = rightLeftNode

	if treeNode.parent != nil {
		if treeNode.parent.left == treeNode {
			treeNode.parent.left = rightNode
		} else {
			treeNode.parent.right = rightNode
		}
	}

	rightNode.parent = treeNode.parent
	treeNode.parent = rightNode
	if rightLeftNode != nil {
		rightLeftNode.parent = treeNode
	}
	/*return sempre da raiz*/
	return rightNode
}

//todo remove

// Method returning number of elements in tree
func (tree RBTree) Count() int {
	return tree.count
}

// Method to print Red-Black tree contents
func (tree RBTree) PrintTree() {
	printRBTreeUtil(tree.root)
}

// Method to print Red-Black tree layout
// Passed argument defines how much spacing there is between tree levels
func (tree RBTree) PrintTree2D(spacing int) {
	// Pass initial space count as 0
	printRBTree2DUtil(tree.root, 0, spacing)
}

// General purpose binary tree printing utility function
func printRBTreeUtil(tree *treeNode) {
	if tree == nil {
		return
	}
	printRBTreeUtil(tree.left)
	if tree.isRed {
		fmt.Println(tree.key, tree.cargo, "R")
	} else {
		fmt.Println(tree.key, tree.cargo, "B")
	}
	printRBTreeUtil(tree.right)
}

// General purpose 2D printing utility function
func printRBTree2DUtil(tree *treeNode, space int, spacing int) {

	// Base case
	if tree == nil {
		return
	}

	// Increase distance between levels
	space += spacing

	// Process right child first
	printRBTree2DUtil(tree.right, space, spacing)

	// Print current node after space
	// count
	for i := spacing; i < space; i++ {
		fmt.Print(" ")
	}

	if tree.isRed {
		fmt.Println(tree.cargo, "R")
	} else {
		fmt.Println(tree.cargo, "B")
	}

	// Process left child

	printRBTree2DUtil(tree.left, space, spacing)
}
