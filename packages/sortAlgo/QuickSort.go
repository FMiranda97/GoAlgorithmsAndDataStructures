package sortAlgo

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"sync"
)

// Function to start quick sort in slice
func QuickSort(arr interface{}) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = errors.New("sort failed. " + fmt.Sprintf("%v", r))
		}
	}()
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

// Utility function to swap elements in quick sort
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

// Function to start quick sort in slice using concurrency
func QuickSortC(arr interface{}) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = errors.New("sort failed. " + fmt.Sprintf("%v", r))
		}
	}()
	if slice, swap, err := sortSetup(arr); err == nil {
		quickSortC(0, slice.Len()-1, slice, swap)
		return err
	} else {
		return err
	}
}

// Utility function to perform quick sort in slice
func quickSortC(l int, r int, slice reflect.Value, swap func(int, int)) {
	if l < r {
		m := partition(l, r, slice, swap)
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			quickSortC(l, m-1, slice, swap)
			wg.Done()
		}()
		go func() {
			quickSortC(m+1, r, slice, swap)
			wg.Done()
		}()
		wg.Wait()
	}
}
