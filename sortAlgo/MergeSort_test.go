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

/*func TestMergeSortPerformance(t *testing.T) {
	var p []pessoa
	startB := time.Now()
	_ = BubbleSort(sortedArray[:])
	endB := time.Since(startB)
	copy(p[:], randomArray[:])
	startSM := time.Now()
	_ = MergeSort(p[:])
	endSM := time.Since(startSM)
	copy(p[:], randomArray[:])
	startCM := time.Now()
	_ = MergeSortConcurrent(p[:])
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
}*/
