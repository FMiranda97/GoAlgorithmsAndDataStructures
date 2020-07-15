package ADS

import (
	"errors"
	"reflect"
)

// Splay tree object
type SplayTree struct {
	root  *treeNode
	count int
}

// Splay tree constructor
func NewSplayTree() SplayTree {
	return SplayTree{
		root:  nil,
		count: 0,
	}
}

// Method to insert cargo into splay tree
func (tree *SplayTree) Insert(cargo Sortable) (err error) {
	defer catch(&err)
	if tree.count > 0 && reflect.TypeOf(tree.root.cargo) != reflect.TypeOf(cargo) {
		return errors.New("inserted cargo not of same type as previously inserted cargo")
	}
	newNode := &treeNode{
		cargo: cargo,
	}
	tree.root, err = insertSplayTreeUtil(tree.root, newNode)
	if err == nil {
		tree.count++
	}
	tree.root, err = splay(newNode)
	return
}

// Utility function for splay tree insertion
func insertSplayTreeUtil(tree *treeNode, newNode *treeNode) (*treeNode, error) {
	var err error
	if tree == nil {
		return newNode, nil
	}

	if newNode.cargo.CompareTo(tree.cargo) < 0 {
		tree.left, err = insertSplayTreeUtil(tree.left, newNode)
		tree.left.parent = tree
	} else if newNode.cargo.CompareTo(tree.cargo) > 0 {
		tree.right, err = insertSplayTreeUtil(tree.right, newNode)
		tree.right.parent = tree
	} else {
		return tree, errors.New("element with this key already exists")
	}
	return tree, err
}

// Function to bring given node to root of tree
func splay(tree *treeNode) (*treeNode, error) {
	var err error
	for tree.parent != nil {
		if tree.parent.parent == nil {
			if tree == tree.parent.left {
				rightRotateSplay(tree.parent)
			} else if tree == tree.parent.right {
				leftRotateSplay(tree.parent)
			}
		} else if tree == tree.parent.right && tree.parent == tree.parent.parent.right {
			leftRotateSplay(tree.parent.parent)
			leftRotateSplay(tree.parent)
		} else if tree == tree.parent.left && tree.parent == tree.parent.parent.left {
			rightRotateSplay(tree.parent.parent)
			rightRotateSplay(tree.parent)
		} else if tree == tree.parent.left && tree.parent == tree.parent.parent.right {
			rightRotateSplay(tree.parent)
			leftRotateSplay(tree.parent)
		} else if tree == tree.parent.right && tree.parent == tree.parent.parent.left {
			leftRotateSplay(tree.parent)
			rightRotateSplay(tree.parent)
		}
	}
	return tree, err
}

// Utility function to right rotate subtree rooted with tree for splay trees
func rightRotateSplay(tree *treeNode) *treeNode {
	leftNode := tree.left
	leftRightNode := leftNode.right

	leftNode.right = tree
	tree.left = leftRightNode

	if tree.parent != nil {
		if tree.parent.right == tree {
			tree.parent.right = leftNode
		} else {
			tree.parent.left = leftNode
		}
	}

	leftNode.parent = tree.parent
	tree.parent = leftNode

	if leftRightNode != nil {
		leftRightNode.parent = tree
	}

	return leftNode
}

// Utility function to left rotate subtree rooted with tree for splay trees
func leftRotateSplay(tree *treeNode) *treeNode {
	rightNode := tree.right
	rightLeftNode := rightNode.left

	rightNode.left = tree
	tree.right = rightLeftNode

	if tree.parent != nil {
		if tree.parent.left == tree {
			tree.parent.left = rightNode
		} else {
			tree.parent.right = rightNode
		}
	}

	rightNode.parent = tree.parent
	tree.parent = rightNode

	if rightLeftNode != nil {
		rightLeftNode.parent = tree
	}

	return rightNode
}

// Method to retrieve cargo from splay tree
func (tree *SplayTree) Get(cargo Sortable) (_ Sortable, e error) {
	defer catch(&e)
	found, err := getBinaryTreeUtil(tree.root, cargo)
	if err == nil {
		tree.root, err = splay(found)
		return found.cargo, err
	}
	return nil, err
}

// Method returning number of elements in tree
func (tree SplayTree) Count() int {
	return tree.count
}

// Method to print splay tree contents
func (tree SplayTree) Print() {
	printTreeUtil(tree.root)
}

// Method to print splay tree layout
// Passed argument defines how much spacing there is between tree levels
func (tree SplayTree) PrintTree2D(spacing int) {
	// Pass initial space count as 0
	printTree2DUtil(tree.root, 0, spacing)
}

// Method to remove cargo from a splay tree
func (tree *SplayTree) Remove(cargo Sortable) (e error) {
	defer catch(&e)
	_, e = tree.Get(cargo)
	if e == nil {
		tree.root, e = removeSplayTreeUtil(tree.root)
		tree.count--
	}
	return e
}

// Utility function for splay tree cargo removal
func removeSplayTreeUtil(tree *treeNode) (*treeNode, error) {
	var err error
	if tree.left == nil {
		tree = tree.right
		tree.parent = nil
	} else {
		aux := tree
		var rightmost *treeNode
		for rightmost = tree.left; rightmost.right != nil; rightmost = rightmost.right {
		}
		tree, err = splay(rightmost)
		tree.right = aux.right
		aux.right.parent = tree
	}
	return tree, err
}
