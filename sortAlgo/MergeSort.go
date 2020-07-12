package sortAlgo

import "reflect"

//todo get a better concurrent method

// Function to start merge sort in array
func MergeSort(arr interface{}) error { // why does this work
	if slice, _, err := sortSetup(arr); err == nil {
		mergeSort(0, slice.Len()-1, slice)
		return err
	} else {
		return err
	}
}

// Utility function to perform merge sort in array
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

	/* Copy data to temp arrays */
	for i := 0; i < n1; i++ {
		left.Index(i).Set(reflect.ValueOf(get(l+i, slice)))
	}
	for i := 0; i < n2; i++ {
		right.Index(i).Set(reflect.ValueOf(get(m+1+i, slice)))
	}
	/* Merge arrays */
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

// Function to start merge sort in array using concurrency
func MergeSortConcurrent(arr interface{}) (err error) { // why does this work
	if slice, _, err := sortSetup(arr); err == nil {
		final := make(chan reflect.Value)
		defer close(final)
		go mergeSortConcurrent(slice, final)
		res := <-final
		for i := 0; i < slice.Len(); i++ {
			slice.Index(i).Set(reflect.ValueOf(get(i, res)))
		}
		return err
	} else {
		return err
	}
}

// Utility function to perform merge sort in array using concurrency
func mergeSortConcurrent(slice reflect.Value, res chan reflect.Value) {
	if slice.Len() == 1 {
		res <- slice
	} else {
		m := slice.Len() / 2
		var left, right reflect.Value
		leftC, rightC := make(chan reflect.Value), make(chan reflect.Value)
		defer close(leftC)
		defer close(rightC)
		leftSlc := slice.Slice(0, m)
		go mergeSortConcurrent(leftSlc, leftC)
		rightSlc := slice.Slice(m, slice.Len())
		go mergeSortConcurrent(rightSlc, rightC)
		left = <-leftC
		right = <-rightC
		res <- mergeConcurrent(left, right)
	}
}

// Utility function to perform merge in merge sort process using concurrency
func mergeConcurrent(left reflect.Value, right reflect.Value) reflect.Value {
	result := reflect.MakeSlice(left.Type(), left.Len()+right.Len(), left.Len()+right.Len()) // todo check type
	l, r, i := 0, 0, 0
	for i = 0; l < left.Len() && r < right.Len(); i++ {
		if get(l, left).CompareTo(get(r, right)) < 0 {
			result.Index(i).Set(reflect.ValueOf(get(l, left)))
			l++
		} else {
			result.Index(i).Set(reflect.ValueOf(get(r, right)))
			r++
		}
	}
	/* finish merging */
	for l < left.Len() {
		result.Index(i).Set(reflect.ValueOf(get(l, left)))
		i++
		l++
	}
	for r < right.Len() {
		result.Index(i).Set(reflect.ValueOf(get(r, right)))
		i++
		r++
	}
	return result
}
