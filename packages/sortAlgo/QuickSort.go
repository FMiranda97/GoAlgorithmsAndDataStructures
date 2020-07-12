package sortAlgo

import (
	"math/rand"
	"reflect"
	"sync"
)

//todo concurrent and dual pivot

// Function to start quick sort in slice
func QuickSort(arr interface{}) error { // why does this work
	if slice, swap, err := sortSetup(arr); err == nil {
		quickSort(0, slice.Len()-1, slice, swap)
		return err
	} else {
		return err
	}
}

// Utility function to perform quick sort in slice
func quickSort(l int, r int, slice reflect.Value, swap func(int, int)) {
	if l < r {
		m := partition(l, r, slice, swap)
		quickSort(l, m-1, slice, swap)
		quickSort(m+1, r, slice, swap)
	}
}

func partition(l int, r int, slice reflect.Value, swap func(int, int)) int {
	swap(l+rand.Int()%(r-l), r) //random pivot
	pivot := get(r, slice)
	for j := l; j < r; j++ {
		if get(j, slice).CompareTo(pivot) < 0 {
			swap(l, j)
			l++
		}
	}
	swap(l, r)
	return l
}

// Function to start merge sort in slice
func QuickSortConcurrent(arr interface{}) error { // why does this work
	if slice, swap, err := sortSetup(arr); err == nil {
		quickSortConcurrent(0, slice.Len()-1, slice, swap)
		return err
	} else {
		return err
	}
}

// Utility function to perform merge sort in slice
func quickSortConcurrent(l int, r int, slice reflect.Value, swap func(int, int)) {
	if l < r {
		m := partitionConcurrent(l, r, slice, swap)
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			quickSort(l, m-1, slice, swap)
			wg.Done()
		}()
		go func() {
			quickSort(m+1, r, slice, swap)
			wg.Done()
		}()
		wg.Wait()
	}
}

func partitionConcurrent(l int, r int, slice reflect.Value, swap func(int, int)) int {
	swap(l+rand.Int()%(r-l), r) //random pivot
	pivot := get(r, slice)
	for j := l; j < r; j++ {
		if get(j, slice).CompareTo(pivot) < 0 {
			swap(l, j)
			l++
		}
	}
	swap(l, r)
	return l
}
