package sortAlgo

import (
	"math/rand"
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

const size = 400_000

var sortedArray, randomArray [size]pessoa
var invalidArray [size]int

func setup() {
	for i := 0; i < size; i++ {
		sortedArray[i] = pessoa{
			id: rand.Int() % (size * 100),
		}
	}
	copy(randomArray[:], sortedArray[:])
	_ = MergeSort(sortedArray[:])
	return
}
