package sortAlgo

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	err := MergeSort(invalidArray[:])
	if err == nil {
		t.Errorf("did not report error on invalid input")
	}
	var p [size]pessoa
	copy(p[:], randomArray[:])
	_ = MergeSort(p[:])
	if !reflect.DeepEqual(sortedArray, p) {
		t.Errorf("sort incorrect")
	}
}

func TestConcurrentMergeSort(t *testing.T) {
	err := MergeSortConcurrent(invalidArray[:])
	if err == nil {
		t.Errorf("did not report error on invalid input")
	}
	var p [size]pessoa
	copy(p[:], randomArray[:])
	_ = MergeSortConcurrent(p[:])
	if !reflect.DeepEqual(sortedArray, p) {
		t.Errorf("sort incorrect")
	}
}

func TestMergeSortC(t *testing.T) {
	err := MergeSortC(invalidArray[:])
	if err == nil {
		t.Errorf("did not report error on invalid input")
	}
	var p [size]pessoa
	copy(p[:], randomArray[:])
	_ = MergeSortC(p[:])
	if !reflect.DeepEqual(sortedArray, p) {
		t.Errorf("sort incorrect")
	}
}
