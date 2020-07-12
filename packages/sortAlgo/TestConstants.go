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

const sizeEfficient = 400_000
const sizeSimple = 1000

var sortedArray, randomArray [sizeEfficient]pessoa
var invalidArray [sizeEfficient]int

func setup() {
	for i := 0; i < sizeEfficient; i++ {
		sortedArray[i] = pessoa{
			id: rand.Int() % (sizeEfficient * 100),
		}
	}
	copy(randomArray[:], sortedArray[:])
	_ = MergeSort(sortedArray[:])
	return
}
