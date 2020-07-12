package sortAlgo

import (
	"reflect"
	"testing"
)

func TestQuickSortDual(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	err := QuickSortDual(invalidArray[:])
	if err == nil {
		t.Errorf("did not report error on invalid input")
	}
	var p [sizeEfficient]pessoa
	copy(p[:], randomArray[:])
	_ = QuickSortDual(p[:])
	if !reflect.DeepEqual(sortedArray, p) {
		t.Errorf("sort incorrect")
	}
}

func TestQuickSortDualConcurrent(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	err := QuickSortDualC(invalidArray[:])
	if err == nil {
		t.Errorf("did not report error on invalid input")
	}
	var p [sizeEfficient]pessoa
	copy(p[:], randomArray[:])
	_ = QuickSortDualC(p[:])
	if !reflect.DeepEqual(sortedArray, p) {
		t.Errorf("sort incorrect")
	}
}
