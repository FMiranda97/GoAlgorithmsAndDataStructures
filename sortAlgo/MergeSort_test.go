package sortAlgo

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
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

const size = 4_000

var p1, p2, p3 [size]pessoa
var invalidArray [size]int

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func setup() {
	for i := 0; i < size; i++ {
		p1[i] = pessoa{
			id: rand.Int() % (size * 100),
		}
	}
	copy(p2[:], p1[:])
	copy(p3[:], p1[:])
	_ = BubbleSort(p1[:])
	return
}

func TestMergeSort(t *testing.T) {
	err := MergeSort(invalidArray[:])
	if err == nil {
		t.Errorf("did not report error on invalid input")
	}
	var p [size]pessoa
	copy(p[:], p2[:])
	_ = MergeSort(p[:])
	if !reflect.DeepEqual(p1, p) {
		t.Errorf("sort incorrect")
	}
}

func TestConcurrentMergeSort(t *testing.T) {
	err := MergeSortConcurrent(invalidArray[:])
	if err == nil {
		t.Errorf("did not report error on invalid input")
	}
	var p [size]pessoa
	copy(p[:], p2[:])
	_ = MergeSortConcurrent(p[:])
	if !reflect.DeepEqual(p1, p) {
		t.Errorf("sort incorrect")
	}
}

func TestMergeSortPerformance(t *testing.T) {
	startB := time.Now()
	_ = BubbleSort(p1[:])
	endB := time.Since(startB)
	startSM := time.Now()
	_ = MergeSort(p2[:])
	endSM := time.Since(startSM)
	startCM := time.Now()
	_ = MergeSortConcurrent(p3[:])
	endCM := time.Since(startCM)
	if endB < endSM {
		t.Errorf("bubble performed better than merge")
		fmt.Println("bubble sort:", endB)
		fmt.Println("simple merge sort:", endSM)
		fmt.Println("concurrent merge sort:", endCM)
	}
	if endSM < endCM {
		t.Errorf("simple performed better than concurrent")
		fmt.Println("bubble sort:", endB)
		fmt.Println("simple merge sort:", endSM)
		fmt.Println("concurrent merge sort:", endCM)
	}
}
