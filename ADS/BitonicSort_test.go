package ADS

import (
	"reflect"
	"testing"
)

func TestBitonicSort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	err := BitonicSort(invalidArray[:])
	if err == nil {
		t.Errorf("did not report error on invalid input")
	}
	var p [sizeEfficient]pessoa
	copy(p[:], randomArray[:])
	_ = BitonicSort(p[:])
	if !reflect.DeepEqual(sortedArray, p) {
		t.Errorf("sort incorrect")
	}
}

func TestConcurrentBitonicSort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	err := BitonicSortC(invalidArray[:])
	if err == nil {
		t.Errorf("did not report error on invalid input")
	}
	var p [sizeEfficient]pessoa
	copy(p[:], randomArray[:])
	_ = BitonicSortC(p[:])
	if !reflect.DeepEqual(sortedArray, p) {
		t.Errorf("sort incorrect")
	}
}
