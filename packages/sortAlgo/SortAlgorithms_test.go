package sortAlgo

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestSimpleSorts(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	var p1, p2, p3 [sizeSimple]pessoa
	for i := 0; i < sizeSimple; i++ {
		p1[i] = pessoa{
			id: rand.Int() % (sizeSimple * 100),
		}
	}
	copy(p2[:], p1[:])
	copy(p3[:], p1[:])
	_ = BubbleSort(p1[:])
	_ = InsertionSort(p2[:])
	_ = SelectionSort(p3[:])
	if !reflect.DeepEqual(p1, p2) {
		t.Errorf("insertion sort failed")
	}
	if !reflect.DeepEqual(p1, p3) {
		t.Errorf("bubble sort failed")
	}
}

func TestSortPerformance(t *testing.T) {
	var p [sizeEfficient]pessoa

	copy(p[:], randomArray[:])
	startBitonicC := time.Now()
	_ = BitonicSortC(p[:])
	endBitonicC := time.Since(startBitonicC)
	fmt.Println("concurrent bitonic sort:", endBitonicC)

	copy(p[:], randomArray[:])
	startBitonic := time.Now()
	_ = BitonicSort(p[:])
	endBitonic := time.Since(startBitonic)
	fmt.Println("bitonic sort:", endBitonic)

	copy(p[:], randomArray[:])
	startQuickDualC := time.Now()
	_ = QuickSortDualC(p[:])
	endQuickDualC := time.Since(startQuickDualC)
	fmt.Println("concurrent quick sort dual:", endQuickDualC)

	copy(p[:], randomArray[:])
	startQuickDual := time.Now()
	_ = QuickSortDual(p[:])
	endQuickDual := time.Since(startQuickDual)
	fmt.Println("quick sort dual:", endQuickDual)

	copy(p[:], randomArray[:])
	startQuickC := time.Now()
	_ = QuickSortC(p[:])
	endQuickC := time.Since(startQuickC)
	fmt.Println("concurrent quick sort:", endQuickC)

	copy(p[:], randomArray[:])
	startQuick := time.Now()
	_ = QuickSort(p[:])
	endQuick := time.Since(startQuick)
	fmt.Println("quick sort:", endQuick)

	copy(p[:], randomArray[:])
	startMergeC := time.Now()
	_ = MergeSortC(p[:])
	endMergeC := time.Since(startMergeC)
	fmt.Println("concurrent merge sort:", endMergeC)

	copy(p[:], randomArray[:])
	startMerge := time.Now()
	_ = MergeSort(p[:])
	endMerge := time.Since(startMerge)
	fmt.Println("simple merge sort:", endMerge)

	if false {
		t.Errorf("wtf")
	}
}
