package main

import (
	. "./packages/sortAlgo"
	"fmt"
	"math/rand"
	"time"
)

type pessoa struct {
	id int
}

func (p pessoa) CompareTo(t interface{}) int8 {
	target := t.(pessoa)
	if p.id < target.id {
		return -1
	} else if p.id > target.id {
		return 1
	} else {
		return 0
	}
}

func main() {
	const size = 10_000_000
	var arr [size]pessoa
	for i := 0; i < size; i++ {
		arr[i] = pessoa{
			id: rand.Int() % (size * 100),
		}
	}
	start := time.Now()
	_ = Sort(arr[:])
	fmt.Println("sort time:", time.Since(start))
}
