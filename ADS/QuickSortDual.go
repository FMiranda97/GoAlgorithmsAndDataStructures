package ADS

import (
	"math/rand"
	"reflect"
	"sync"
)

// Function to start merge sort in slice using 2 pivots
func QuickSortDual(arr interface{}) (e error) {
	defer panicControl(&e)
	if slice, swap, err := sortSetup(arr); err == nil {
		quickSortDual(0, slice.Len()-1, slice, swap)
		return err
	} else {
		return err
	}
}

// Utility function to perform merge sort in slice
func quickSortDual(l int, r int, slice reflect.Value, swap func(int, int)) {
	if l < r {
		lp, rp := partitionDual(l, r, slice, swap)
		quickSortDual(l, lp-1, slice, swap)
		quickSortDual(lp+1, rp-1, slice, swap)
		quickSortDual(rp+1, r, slice, swap)
	}
}

// Utility function to swap elements in quick sort dual
func partitionDual(low int, high int, slice reflect.Value, swap func(int, int)) (int, int) {
	swap(low, low+rand.Int()%(high-low))
	swap(high, low+rand.Int()%(high-low))
	if get(low, slice).CompareTo(get(high, slice)) > 0 {
		swap(low, high)
	}
	j := low + 1
	g := high - 1
	k := low + 1
	leftPivot := get(low, slice)
	rightPivot := get(high, slice)
	for k <= g {
		if get(k, slice).CompareTo(leftPivot) < 0 {
			swap(k, j)
			j++
		} else if get(k, slice).CompareTo(rightPivot) >= 0 {
			for get(g, slice).CompareTo(rightPivot) > 0 && k < g {
				g--
			}
			swap(k, g)
			g--
			if get(k, slice).CompareTo(leftPivot) < 0 {
				swap(k, j)
				j++
			}
		}
		k++
	}
	j--
	g++
	swap(low, j)
	swap(high, g)
	return j, g
}

// Function to start merge sort in slice using 2 pivots
func QuickSortDualC(arr interface{}) (e error) {
	defer panicControl(&e)
	if slice, swap, err := sortSetup(arr); err == nil {
		quickSortDualC(0, slice.Len()-1, slice, swap)
		return err
	} else {
		return err
	}
}

// Utility function to perform merge sort in slice
func quickSortDualC(l int, r int, slice reflect.Value, swap func(int, int)) {
	if l < r {
		lp, rp := partitionDual(l, r, slice, swap)
		var wg sync.WaitGroup
		wg.Add(3)
		go func() {
			quickSortDualC(l, lp-1, slice, swap)
			wg.Done()
		}()
		go func() {
			quickSortDualC(lp+1, rp-1, slice, swap)
			wg.Done()
		}()
		go func() {
			quickSortDualC(rp+1, r, slice, swap)
			wg.Done()
		}()
		wg.Wait()
	}
}
