package sortAlgo

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	err := MergeSort(invalidArray[:])
	if err == nil {
		t.Errorf("did not report error on invalid input")
	}
	var p [sizeEfficient]pessoa
	copy(p[:], randomArray[:])
	_ = MergeSort(p[:])
	if !reflect.DeepEqual(sortedArray, p) {
		t.Errorf("sort incorrect")
	}
}

func TestConcurrentMergeSort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	err := MergeSortC(invalidArray[:])
	if err == nil {
		t.Errorf("did not report error on invalid input")
	}
	var p [sizeEfficient]pessoa
	copy(p[:], randomArray[:])
	_ = MergeSortC(p[:])
	if !reflect.DeepEqual(sortedArray, p) {
		t.Errorf("sort incorrect")
	}
}
