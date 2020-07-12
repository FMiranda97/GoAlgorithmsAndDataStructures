package main

import (
	. "./sortAlgo"
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
	const size = 16
	var p1, p2 [size]pessoa
	for i := 0; i < size; i++ {
		p1[i] = pessoa{
			id: rand.Int() % (size * 100),
		}
	}
	copy(p2[:], p1[:])
	var p [size]int
	err := BubbleSort(p[:])
	fmt.Println(err)
	start := time.Now()
	_ = MergeSortConcurrent(p2[:])
	fmt.Println("Merge sort time:", time.Since(start))
}
