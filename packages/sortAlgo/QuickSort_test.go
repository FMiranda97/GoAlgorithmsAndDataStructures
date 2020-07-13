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
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	err := QuickSort(invalidArray[:])
	if err == nil {
		t.Errorf("did not report error on invalid input")
	}
	var p [sizeEfficient]pessoa
	copy(p[:], randomArray[:])
	_ = QuickSort(p[:])
	if !reflect.DeepEqual(sortedArray, p) {
		t.Errorf("sort incorrect")
	}
}

func TestQuickSortConcurrent(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	err := QuickSortC(invalidArray[:])
	if err == nil {
		t.Errorf("did not report error on invalid input")
	}
	var p [sizeEfficient]pessoa
	copy(p[:], randomArray[:])
	_ = QuickSortC(p[:])
	if !reflect.DeepEqual(sortedArray, p) {
		t.Errorf("sort incorrect")
	}
}
