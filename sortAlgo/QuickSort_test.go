package sortAlgo

import (
	"fmt"
	"reflect"
	"testing"
	"time"
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
	_ = BubbleSort(sortedArray)
	_ = QuickSort(p[:])
	fmt.Println(p)
	fmt.Println(randomArray)
	fmt.Println(sortedArray)
	if !reflect.DeepEqual(sortedArray, p) {
		t.Errorf("sort incorrect")
	}
}

func TestQuickSortPerformance(t *testing.T) {
	startB := time.Now()
	_ = BubbleSort(sortedArray[:])
	endB := time.Since(startB)
	startQS := time.Now()
	_ = QuickSort(randomArray[:])
	endQS := time.Since(startQS)
	if endB < endQS {
		t.Errorf("bubble performed better than quick")
		fmt.Println("bubble sort:", endB)
		fmt.Println("simple merge sort:", endQS)
	}
}
