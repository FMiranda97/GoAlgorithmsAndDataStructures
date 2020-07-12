package sortAlgo

import (
	"fmt"
	"testing"
	"time"
)

func TestSortPerformance(t *testing.T) {
	var p [size]pessoa

	copy(p[:], randomArray[:])
	startQuick := time.Now()
	_ = QuickSort(p[:])
	endQuick := time.Since(startQuick)

	copy(p[:], randomArray[:])
	startQuickC := time.Now()
	_ = QuickSortConcurrent(p[:])
	endQuickC := time.Since(startQuickC)

	copy(p[:], randomArray[:])
	startMerge := time.Now()
	_ = MergeSort(p[:])
	endMerge := time.Since(startMerge)

	copy(p[:], randomArray[:])
	startMergeC := time.Now()
	_ = MergeSortC(p[:])
	endMergeC := time.Since(startMergeC)

	if false {
		t.Errorf("wtf")
	}
	fmt.Println("quick sort:", endQuick)
	fmt.Println("concurrent quick sort:", endQuickC)
	fmt.Println("simple merge sort:", endMerge)
	fmt.Println("concurrent merge sort:", endMergeC)
}
