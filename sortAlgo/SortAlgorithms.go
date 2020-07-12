package sortAlgo

import (
	"errors"
	"reflect"
)

// interface that must be implemented by elements of slices using this package
// CompareTo must return -1, 0 or 1 if target is less than argument, equal or greater respectively
type Sortable interface {
	CompareTo(interface{}) int8
}

// utility function to check if data types are correct and returning function to swap elements
func sortSetup(arr interface{}) (reflect.Value, func(int, int), error) {
	if reflect.TypeOf(arr).Kind() == reflect.Slice {
		slice := reflect.ValueOf(arr)
		swap := reflect.Swapper(slice.Interface())
		return slice, swap, nil
	}
	return reflect.Value{}, nil, errors.New("argument is not pointer to slice or does not implement required interface")
}

// utility function to return generic element in slice
func get(index int, t reflect.Value) Sortable {
	return t.Index(index).Interface().(Sortable)
}

// function to start bubble sort in array
func BubbleSort(arr interface{}) (err error) { // why does this work
	if slice, swap, err := sortSetup(arr); err == nil {
		for i := 0; i < slice.Len()-1; i++ {
			for j := 0; j < slice.Len()-1; j++ {
				a, b := get(j, slice), get(j+1, slice)
				if a.CompareTo(b) > 0 {
					swap(j, j+1)
				}
			}
		}
	}
	return err
}
