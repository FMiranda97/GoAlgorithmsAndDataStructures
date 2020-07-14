package ADS

//TODO shell, heap

import (
	"errors"
	"reflect"
)

// Utility function to check if data types are correct and returning function to swap elements
func sortSetup(arr interface{}) (reflect.Value, func(int, int), error) {
	typ := reflect.TypeOf((*Sortable)(nil)).Elem()
	if reflect.TypeOf(arr).Kind() == reflect.Slice && reflect.TypeOf(arr).Elem().Implements(typ) {
		slice := reflect.ValueOf(arr)
		swap := reflect.Swapper(slice.Interface())
		return slice, swap, nil
	}
	return reflect.Value{}, nil, errors.New("argument is not pointer to slice or does not implement required interface")
}

// Utility function to return generic element in slice
func get(index int, t reflect.Value) Sortable {
	return t.Index(index).Interface().(Sortable)
}

// Function to start bubble sort in slice
func BubbleSort(arr interface{}) (e error) {
	defer catch(&e)
	if slice, swap, err := sortSetup(arr); err == nil {
		for i := 0; i < slice.Len()-1; i++ {
			for j := 0; j < slice.Len()-1; j++ {
				a, b := get(j, slice), get(j+1, slice)
				if a.CompareTo(b) > 0 {
					swap(j, j+1)
				}
			}
		}
		return nil
	} else {
		return err
	}
}

// Function to start insertion sort in slice
func InsertionSort(arr interface{}) (e error) {
	defer catch(&e)
	if slice, _, err := sortSetup(arr); err == nil {
		for i := 1; i < slice.Len(); i++ {
			key := get(i, slice)
			j := i - 1
			for j >= 0 && get(j, slice).CompareTo(key) > 0 {
				slice.Index(j + 1).Set(slice.Index(j))
				j = j - 1
			}
			slice.Index(j + 1).Set(reflect.ValueOf(key))
		}
		return nil
	} else {
		return err
	}
}

// Function to start selection sort in slice
func SelectionSort(arr interface{}) (e error) {
	defer catch(&e)
	if slice, swap, err := sortSetup(arr); err == nil {
		for i := 0; i < slice.Len()-1; i++ {
			min := i
			for j := i + 1; j < slice.Len(); j++ {
				if get(j, slice).CompareTo(get(min, slice)) < 0 {
					min = j
				}
			}
			if get(i, slice).CompareTo(get(min, slice)) != 0 {
				swap(i, min)
			}
		}
		return nil
	} else {
		return err
	}
}

// Function to start shell sort in slice
func ShellSort(arr interface{}) (e error) {
	defer catch(&e)
	if slice, _, err := sortSetup(arr); err == nil {
		gapSize := slice.Len() / 2
		for ; gapSize > 0; gapSize /= 2 {
			for current := gapSize; current < slice.Len(); current++ {
				valueSort := get(current, slice)
				valuesComp := current - gapSize
				for ; valuesComp >= 0 && valueSort.CompareTo(get(valuesComp, slice)) < 0; valuesComp -= gapSize {
					slice.Index(valuesComp + gapSize).Set(slice.Index(valuesComp))
				}
				slice.Index(valuesComp + gapSize).Set(reflect.ValueOf(valueSort))
			}
		}
		return nil
	} else {
		return err
	}
}

// Default function to sort in slice
func Sort(arr interface{}) error {
	return QuickSortDualC(arr)
}
