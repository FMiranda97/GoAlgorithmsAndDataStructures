package linkedList

type linkedListCargoString interface {
	getKey() string
}

type linkedListCargoInt64 interface {
	getKey() int64
}

type linkedListCargoInt interface {
	getKey() int
}

type LinkedList = *node

type node struct {
	cargo interface{}
	next  *node
}

func implementsLinkedListCargo(cargo interface{}) bool {
	if _, ok := cargo.(linkedListCargoString); ok {
		return true
	}
	if _, ok := cargo.(linkedListCargoInt); ok {
		return true
	}
	if _, ok := cargo.(linkedListCargoInt64); ok {
		return true
	}
	return false
}
