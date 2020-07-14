package ADS

import (
	"reflect"
	"sync"
)

// Function to start quick sort in slice
func BitonicSort(arr interface{}) (e error) {
	defer panicControl(&e)
	if slice, swap, err := sortSetup(arr); err == nil {
		bitonicSort(0, slice.Len(), slice, swap, true)
		return err
	} else {
		return err
	}
}

// Utility function to perform quick sort in slice
func bitonicSort(l int, cnt int, slice reflect.Value, swap func(int, int), up bool) {
	if cnt > 1 {
		k := cnt / 2
		bitonicSort(l, k, slice, swap, !up)
		bitonicSort(l+k, cnt-k, slice, swap, up)
		bitonicMerge(l, cnt, slice, swap, up)
	}
}

// Utility function to swap elements in quick sort
func bitonicMerge(l int, cnt int, slice reflect.Value, swap func(int, int), up bool) {
	if cnt > 1 {
		k := greatestPowerOfTwoLessThan(cnt)
		for i := l; i < l+cnt-k; i++ {
			if up == (get(i, slice).CompareTo(get(i+k, slice)) > 0) {
				swap(i, i+k)
			}
		}
		bitonicMerge(l, k, slice, swap, up)
		bitonicMerge(l+k, cnt-k, slice, swap, up)
	}
}

func greatestPowerOfTwoLessThan(n int) int {
	k := 1
	for k > 0 && k < n {
		k = k << 1
	}
	return k >> 1
}

// Function to start quick sort in slice
func BitonicSortC(arr interface{}) (e error) {
	defer panicControl(&e)
	if slice, swap, err := sortSetup(arr); err == nil {
		bitonicSortC(0, slice.Len(), slice, swap, true)
		return err
	} else {
		return err
	}
}

// Utility function to perform quick sort in slice
func bitonicSortC(l int, cnt int, slice reflect.Value, swap func(int, int), up bool) {
	if cnt > 1 {
		k := cnt / 2
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			bitonicSortC(l, k, slice, swap, !up)
			wg.Done()
		}()
		go func() {
			bitonicSortC(l+k, cnt-k, slice, swap, up)
			wg.Done()
		}()
		wg.Wait()
		bitonicMergeC(l, cnt, slice, swap, up)
	}
}

// Utility function to swap elements in quick sort
func bitonicMergeC(l int, cnt int, slice reflect.Value, swap func(int, int), up bool) {
	if cnt > 1 {
		k := greatestPowerOfTwoLessThan(cnt)
		for i := l; i < l+cnt-k; i++ {
			if up == (get(i, slice).CompareTo(get(i+k, slice)) > 0) {
				swap(i, i+k)
			}
		}
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			bitonicMergeC(l, k, slice, swap, up)
			wg.Done()
		}()
		go func() {
			bitonicMergeC(l+k, cnt-k, slice, swap, up)
			wg.Done()
		}()
		wg.Wait()
	}
}
