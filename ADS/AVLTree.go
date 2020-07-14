package ADS

import (
	"errors"
	"reflect"
)

// AVL tree object
type AVLTree struct {
	root  *treeNode
	count int
}

// AVL tree constructor
func NewAVLTree() AVLTree {
	return AVLTree{
		root:  nil,
		count: 0,
	}
}

// Method to insert cargo into AVL tree
func (tree *AVLTree) Insert(cargo Sortable) (err error) {
	defer panicControl(&err)
	if tree.count > 0 && reflect.TypeOf(tree.root.cargo) != reflect.TypeOf(cargo) {
		return errors.New("inserted cargo not of same type as previously inserted cargo")
	}
	tree.root, err = insertAVLTreeUtil(tree.root, cargo)
	if err == nil {
		tree.count++
	}
	return
}

// Utility function for AVL tree insertion
func insertAVLTreeUtil(tree *treeNode, cargo Sortable) (*treeNode, error) {
	var err error
	if tree == nil {
		newNode := treeNode{
			cargo:  cargo,
			height: 1,
		}
		return &newNode, nil
	}
	if cargo.CompareTo(tree.cargo) < 0 {
		tree.left, err = insertAVLTreeUtil(tree.left, cargo)
	} else if cargo.CompareTo(tree.cargo) > 0 {
		tree.right, err = insertAVLTreeUtil(tree.right, cargo)
	} else { // Equal keys are not allowed in BST
		return tree, errors.New("element with this key already exists")
	}

	tree.height = getMax(getHeight(tree.left), getHeight(tree.right)) + 1

	diffHeight := getDiffHeight(tree)

	// Left Left Case
	if diffHeight > 1 && tree.left != nil && cargo.CompareTo(tree.left.cargo) < 0 {
		return rightRotateAVL(tree), err
	}

	// Right Right Case
	if diffHeight < -1 && tree.right != nil && cargo.CompareTo(tree.right.cargo) > 0 {
		return leftRotateAVL(tree), err
	}

	// Left Right Case
	if diffHeight > 1 && tree.left != nil && cargo.CompareTo(tree.left.cargo) > 0 {
		tree.left = leftRotateAVL(tree.left)
		return rightRotateAVL(tree), err
	}

	// Right Left Case
	if diffHeight < -1 && tree.right != nil && cargo.CompareTo(tree.right.cargo) < 0 {
		tree.right = rightRotateAVL(tree.right)
		return leftRotateAVL(tree), err
	}
	return tree, err
}

// A utility function to get height of the tree
func getHeight(treeNode *treeNode) int {
	if treeNode == nil {
		return 0
	} else {
		return treeNode.height
	}
}

// Get height difference of nodeOrder treeNode
func getDiffHeight(treeNode *treeNode) int {
	if treeNode == nil {
		return 0
	}
	return getHeight(treeNode.left) - getHeight(treeNode.right)
}

// A utility function to right rotate subtree rooted with treeNode
func rightRotateAVL(treeNode *treeNode) *treeNode {
	leftNode := treeNode.left
	leftRightNode := leftNode.right

	leftNode.right = treeNode
	treeNode.left = leftRightNode

	treeNode.height = getMax(getHeight(treeNode.right), getHeight(treeNode.left)) + 1
	leftNode.height = getMax(getHeight(leftNode.right), getHeight(leftNode.left)) + 1

	return leftNode
}

// A utility function to left rotate subtree rooted with treeNode
func leftRotateAVL(treeNode *treeNode) *treeNode {
	rightNode := treeNode.right
	rightLeftNode := rightNode.left

	rightNode.left = treeNode
	treeNode.right = rightLeftNode

	treeNode.height = getMax(getHeight(treeNode.right), getHeight(treeNode.left)) + 1
	rightNode.height = getMax(getHeight(rightNode.right), getHeight(rightNode.left)) + 1

	return rightNode
}

// Method to remove cargo from an AVL tree
func (tree *AVLTree) Remove(cargo Sortable) (e error) {
	defer panicControl(&e)
	tree.root, e = removeAVLTreeUtil(tree.root, cargo)
	if e == nil {
		tree.count--
	}
	return e
}

// Remove cargo from AVL tree with a given key
func removeAVLTreeUtil(tree *treeNode, cargo Sortable) (*treeNode, error) {
	if tree == nil {
		return tree, errors.New("no element found with given key")
	}
	var err error
	if cargo.CompareTo(tree.cargo) < 0 {
		tree.left, err = removeAVLTreeUtil(tree.left, cargo)
	} else if cargo.CompareTo(tree.cargo) > 0 {
		tree.right, err = removeAVLTreeUtil(tree.right, cargo)
	} else {
		// nodeOrder with only one child or no child
		if tree.left == nil {
			return tree.right, nil
		} else if tree.right == nil {
			return tree.left, nil
		} else { // nodeOrder with two children
			var rightmost *treeNode
			for rightmost = tree.left; rightmost.right != nil; rightmost = rightmost.right {
			}
			tree.cargo = rightmost.cargo
			tree.left, err = removeAVLTreeUtil(tree.left, rightmost.cargo)
		}
	}

	tree.height = 1 + getMax(getHeight(tree.left), getHeight(tree.right))

	diffHeight := getDiffHeight(tree)

	// Left Left Case
	if diffHeight > 1 && getDiffHeight(tree.left) >= 0 {
		return rightRotateAVL(tree), err
	}

	// Left Right Case
	if diffHeight > 1 && getDiffHeight(tree.left) < 0 {
		tree.left = leftRotateAVL(tree.left)
		return rightRotateAVL(tree), err
	}

	// Right Right Case
	if diffHeight < -1 && getDiffHeight(tree.right) <= 0 {
		return leftRotateAVL(tree), err
	}

	// Right Left Case
	if diffHeight < -1 && getDiffHeight(tree.right) > 0 {
		tree.right = rightRotateAVL(tree.right)
		return leftRotateAVL(tree), err
	}
	return tree, err
}

// Method to retrieve cargo with a given key
func (tree AVLTree) Get(cargo Sortable) (_ Sortable, e error) {
	defer panicControl(&e)
	found, err := getBinaryTreeUtil(tree.root, cargo)
	if err == nil {
		return found.cargo, err
	}
	return nil, err
}

// Method returning number of elements in tree
func (tree AVLTree) Count() int {
	return tree.count
}

// Method to print AVL tree contents
func (tree AVLTree) Print() {
	printTreeUtil(tree.root)
}

// Method to print AVL tree layout
// Passed argument defines how much spacing there is between tree levels
func (tree AVLTree) PrintTree2D(spacing int) {
	// Pass initial space count as 0
	printTree2DUtil(tree.root, 0, spacing)
}
