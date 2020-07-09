package linkedList

type cargoLinkedListString interface {
	GetKey() string
}

type cargoLinkedListInt64 interface {
	GetKey() int64
}

type cargoLinkedListInt interface {
	GetKey() int
}

type LinkedList = *node
type node struct {
	cargo interface{}
	next  *node
}

type BinaryTree = *treeNode
type treeNode struct {
	cargo interface{}
	left  *treeNode
	right *treeNode
}

func implementsLinkedListCargo(cargo interface{}) bool {
	if _, ok := cargo.(cargoLinkedListString); ok {
		return true
	}
	if _, ok := cargo.(cargoLinkedListInt); ok {
		return true
	}
	if _, ok := cargo.(cargoLinkedListInt64); ok {
		return true
	}
	return false
}
