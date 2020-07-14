package ADS

import "reflect"

// Function to start heap sort in slice
func HeapSort(arr interface{}) (e error) {
	defer catch(&e)
	if slice, swap, err := sortSetup(arr); err == nil {
		heapSort(slice, swap)
		return err
	} else {
		return err
	}
}

// Utility function to perform heap sort in slice
func heapSort(slice reflect.Value, swap func(int, int)) {
	for i := slice.Len()/2 - 1; i >= 0; i-- {
		heapify(slice, slice.Len(), i, swap)
	}
	for i := slice.Len() - 1; i > 0; i-- {
		swap(0, i)
		heapify(slice, i, 0, swap)
	}
}

// Utility function for heapSort
func heapify(slice reflect.Value, tam int, index int, swap func(int, int)) {
	largest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < tam && get(left, slice).CompareTo(get(largest, slice)) > 0 {
		largest = left
	}
	if right < tam && get(right, slice).CompareTo(get(largest, slice)) > 0 {
		largest = right
	}

	if largest != index {
		swap(index, largest)
		heapify(slice, tam, largest, swap)
	}
}
