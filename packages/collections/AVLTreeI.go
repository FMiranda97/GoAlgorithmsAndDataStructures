package collections

import (
	"errors"
)

// AVL tree object
type AVLTreeI struct {
	root  *treeNodeI
	count int
}

// AVL tree constructor
func NewAVLTreeI() AVLTreeI {
	return AVLTreeI{
		root:  nil,
		count: 0,
	}
}

// Method to insert cargo into AVL tree
func (tree *AVLTreeI) Insert(key int64, cargo interface{}) (err error) {
	tree.root, err = insertAVLTreeIUtil(tree.root, key, cargo)
	if err == nil {
		tree.count++
	}
	return
}

// Utility function for AVL tree insertion
func insertAVLTreeIUtil(tree *treeNodeI, key int64, cargo interface{}) (*treeNodeI, error) {
	var err error
	if tree == nil {
		newNode := treeNodeI{
			key:    key,
			cargo:  cargo,
			height: 1,
		}
		return &newNode, nil
	}
	if key < tree.key {
		tree.left, err = insertAVLTreeIUtil(tree.left, key, cargo)
	} else if key > tree.key {
		tree.right, err = insertAVLTreeIUtil(tree.right, key, cargo)
	} else { // Equal keys are not allowed in BST
		return tree, errors.New("element with this key already exists")
	}

	tree.height = getMax(getHeightI(tree.left), getHeightI(tree.right)) + 1

	diffHeight := getDiffHeightI(tree)

	// Left Left Case
	if diffHeight > 1 && tree.left != nil && key < tree.left.key {
		return rightRotateAVLI(tree), err
	}

	// Right Right Case
	if diffHeight < -1 && tree.right != nil && key > tree.right.key {
		return leftRotateAVLI(tree), err
	}

	// Left Right Case
	if diffHeight > 1 && tree.left != nil && key > tree.left.key {
		tree.left = leftRotateAVLI(tree.left)
		return rightRotateAVLI(tree), err
	}

	// Right Left Case
	if diffHeight < -1 && tree.right != nil && key < tree.right.key {
		tree.right = rightRotateAVLI(tree.right)
		return leftRotateAVLI(tree), err
	}
	return tree, err
}

// A utility function to get height of the tree
func getHeightI(treeNodeI *treeNodeI) int {
	if treeNodeI == nil {
		return 0
	} else {
		return treeNodeI.height
	}
}

// Get height difference of node treeNodeI
func getDiffHeightI(treeNodeI *treeNodeI) int {
	if treeNodeI == nil {
		return 0
	}
	return getHeightI(treeNodeI.left) - getHeightI(treeNodeI.right)
}

// A utility function to right rotate subtree rooted with treeNodeI
func rightRotateAVLI(treeNodeI *treeNodeI) *treeNodeI {
	leftNode := treeNodeI.left
	leftRightNode := leftNode.right

	leftNode.right = treeNodeI
	treeNodeI.left = leftRightNode

	treeNodeI.height = getMax(getHeightI(treeNodeI.right), getHeightI(treeNodeI.left)) + 1
	leftNode.height = getMax(getHeightI(leftNode.right), getHeightI(leftNode.left)) + 1

	return leftNode
}

// A utility function to left rotate subtree rooted with treeNodeI
func leftRotateAVLI(treeNodeI *treeNodeI) *treeNodeI {
	rightNode := treeNodeI.right
	rightLeftNode := rightNode.left

	rightNode.left = treeNodeI
	treeNodeI.right = rightLeftNode

	treeNodeI.height = getMax(getHeightI(treeNodeI.right), getHeightI(treeNodeI.left)) + 1
	rightNode.height = getMax(getHeightI(rightNode.right), getHeightI(rightNode.left)) + 1

	return rightNode
}

// Method to remove cargo from an AVL tree
func (tree *AVLTreeI) Remove(key int64) error {
	var err error
	tree.root, err = removeAVLTreeIUtil(tree.root, key)
	if err == nil {
		tree.count--
	}
	return err
}

// Remove cargo from AVL tree with a given key
func removeAVLTreeIUtil(tree *treeNodeI, key int64) (*treeNodeI, error) {
	if tree == nil {
		return tree, errors.New("no element found with given key")
	}
	var err error
	if key < tree.key {
		tree.left, err = removeAVLTreeIUtil(tree.left, key)
	} else if key > tree.key {
		tree.right, err = removeAVLTreeIUtil(tree.right, key)
	} else {
		// node with only one child or no child
		if tree.left == nil {
			return tree.right, nil
		} else if tree.right == nil {
			return tree.left, nil
		} else { // node with two children
			var rightmost *treeNodeI
			for rightmost = tree.left; rightmost.right != nil; rightmost = rightmost.right {
			}
			tree.key = rightmost.key
			tree.cargo = rightmost.cargo
			tree.left, err = removeAVLTreeIUtil(tree.left, rightmost.key)
		}
	}

	tree.height = 1 + getMax(getHeightI(tree.left), getHeightI(tree.right))

	diffHeight := getDiffHeightI(tree)

	// Left Left Case
	if diffHeight > 1 && getDiffHeightI(tree.left) >= 0 {
		return rightRotateAVLI(tree), err
	}

	// Left Right Case
	if diffHeight > 1 && getDiffHeightI(tree.left) < 0 {
		tree.left = leftRotateAVLI(tree.left)
		return rightRotateAVLI(tree), err
	}

	// Right Right Case
	if diffHeight < -1 && getDiffHeightI(tree.right) <= 0 {
		return leftRotateAVLI(tree), err
	}

	// Right Left Case
	if diffHeight < -1 && getDiffHeightI(tree.right) > 0 {
		tree.right = rightRotateAVLI(tree.right)
		return leftRotateAVLI(tree), err
	}
	return tree, err
}

// Method returning number of elements in tree
func (tree AVLTreeI) Count() int {
	return tree.count
}

// Method to print AVL tree contents
func (tree AVLTreeI) PrintTree() {
	printTreeIUtil(tree.root)
}

// Method to print AVL tree layout
// Passed argument defines how much spacing there is between tree levels
func (tree AVLTreeI) PrintTree2D(spacing int) {
	// Pass initial space count as 0
	printTreeI2DUtil(tree.root, 0, spacing)
}
