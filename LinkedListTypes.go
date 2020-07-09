package main

type linkedListCargoString interface {
	getKey() string
}

type linkedListCargoInt64 interface {
	getKey() int64
}

type linkedListCargoInt interface {
	getKey() int
}

type linkedList = *node

type node struct {
	cargo interface{}
	next  *node
}
