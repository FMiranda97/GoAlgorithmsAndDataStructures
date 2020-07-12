package sortAlgo

import (
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func TestQuickSort(t *testing.T) {
	err := QuickSort(invalidArray[:])
	if err == nil {
		t.Errorf("did not report error on invalid input")
	}
	var p [size]pessoa
	copy(p[:], randomArray[:])
	_ = QuickSort(p[:])
	if !reflect.DeepEqual(sortedArray, p) {
		t.Errorf("sort incorrect")
	}
}

func TestQuickSortConcurrent(t *testing.T) {
	err := QuickSortConcurrent(invalidArray[:])
	if err == nil {
		t.Errorf("did not report error on invalid input")
	}
	var p [size]pessoa
	copy(p[:], randomArray[:])
	_ = QuickSortConcurrent(p[:])
	if !reflect.DeepEqual(sortedArray, p) {
		t.Errorf("sort incorrect")
	}
}
