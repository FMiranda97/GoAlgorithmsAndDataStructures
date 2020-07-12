package sortAlgo

import (
	"reflect"
	"sync"
)

// Function to start merge sort in slice
func MergeSort(arr interface{}) error {
	if slice, _, err := sortSetup(arr); err == nil {
		mergeSort(0, slice.Len()-1, slice)
		return err
	} else {
		return err
	}
}

// Utility function to perform merge sort in slice
func mergeSort(l int, r int, slice reflect.Value) {
	if r != l {
		m := l + (r-l)/2
		mergeSort(l, m, slice)
		mergeSort(m+1, r, slice)
		merge(l, m, r, slice)
	}
}

// Utility function to perform merge in merge sort process
func merge(l int, m int, r int, slice reflect.Value) {
	n1 := m - l + 1
	n2 := r - m
	left := reflect.MakeSlice(slice.Type(), n1, n1)  // slice.Slice(l, m)
	right := reflect.MakeSlice(slice.Type(), n2, n2) //slice.Slice(m+1, r)

	/* Copy data to temp slices */
	for i := 0; i < n1; i++ {
		left.Index(i).Set(reflect.ValueOf(get(l+i, slice)))
	}
	for i := 0; i < n2; i++ {
		right.Index(i).Set(reflect.ValueOf(get(m+1+i, slice)))
	}
	/* Merge slices */
	i, j, k := 0, 0, l
	for i < n1 && j < n2 {
		if get(i, left).CompareTo(get(j, right)) <= 0 {
			slice.Index(k).Set(reflect.ValueOf(get(i, left)))
			i++
		} else {
			slice.Index(k).Set(reflect.ValueOf(get(j, right)))
			j++
		}
		k++
	}
	/* finish merging */
	for i < n1 {
		slice.Index(k).Set(reflect.ValueOf(get(i, left)))
		i++
		k++
	}
	for j < n2 {
		slice.Index(k).Set(reflect.ValueOf(get(j, right)))
		j++
		k++
	}
}

// Function to start merge sort in slice
func MergeSortC(arr interface{}) error {
	if slice, _, err := sortSetup(arr); err == nil {
		mergeSortC(0, slice.Len()-1, slice)
		return err
	} else {
		return err
	}
}

// Utility function to perform merge sort in slice
func mergeSortC(l int, r int, slice reflect.Value) {
	if r != l {
		m := l + (r-l)/2
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			mergeSortC(l, m, slice)
			wg.Done()
		}()
		go func() {
			mergeSortC(m+1, r, slice)
			wg.Done()
		}()
		wg.Wait()
		merge(l, m, r, slice)
	}
}
