package ADS

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	err := HeapSort(invalidArray[:])
	if err == nil {
		t.Errorf("did not report error on invalid input")
	}
	var p [sizeEfficient]pessoa
	copy(p[:], randomArray[:])
	_ = HeapSort(p[:])
	if !reflect.DeepEqual(sortedArray, p) {
		t.Errorf("sort incorrect")
	}
}
