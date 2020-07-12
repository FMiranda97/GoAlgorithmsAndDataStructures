package sortAlgo

import "reflect"

//todo concurrent and dual pivot

// Function to start merge sort in array
func QuickSort(arr interface{}) error { // why does this work
	if slice, swap, err := sortSetup(arr); err == nil {
		quickSort(0, slice.Len()-1, slice, swap)
		return err
	} else {
		return err
	}
}

// Utility function to perform merge sort in array
func quickSort(l int, r int, slice reflect.Value, swap func(int, int)) {
	if l < r {
		m := partition(l, r, slice, swap)
		quickSort(l, m-1, slice, swap)
		quickSort(m+1, r, slice, swap)
	}
}

func partition(l int, r int, slice reflect.Value, swap func(int, int)) int {
	//todo random pivot
	pivot := get(r, slice)
	i := l - 1 //todo check this
	for j := l; j < r; j++ {
		if get(j, slice).CompareTo(pivot) < 0 {
			i++
			swap(i, j)
		}
	}
	swap(i+1, r)
	return i + 1
}
