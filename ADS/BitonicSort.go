package ADS

import (
	"reflect"
	"sync"
)

// Function to start bitonic sort in slice
func BitonicSort(arr interface{}) (e error) {
	defer catch(&e)
	if slice, swap, err := sortSetup(arr); err == nil {
		bitonicSort(0, slice.Len(), slice, swap, true)
		return err
	} else {
		return err
	}
}

// Utility function to perform bitonic sort in slice
func bitonicSort(l int, cnt int, slice reflect.Value, swap func(int, int), up bool) {
	if cnt > 1 {
		k := cnt / 2
		bitonicSort(l, k, slice, swap, !up)
		bitonicSort(l+k, cnt-k, slice, swap, up)
		bitonicMerge(l, cnt, slice, swap, up)
	}
}

// Utility function to perform merge in bitonic sort
func bitonicMerge(l int, cnt int, slice reflect.Value, swap func(int, int), up bool) {
	if cnt > 1 {
		k := sizePartition(cnt)
		for i := l; i < l+cnt-k; i++ {
			if up == (get(i, slice).CompareTo(get(i+k, slice)) > 0) {
				swap(i, i+k)
			}
		}
		bitonicMerge(l, k, slice, swap, up)
		bitonicMerge(l+k, cnt-k, slice, swap, up)
	}
}

// Function to start bitonic sort in slice using concurrency
func BitonicSortC(arr interface{}) (e error) {
	defer catch(&e)
	if slice, swap, err := sortSetup(arr); err == nil {
		bitonicSortC(0, slice.Len(), slice, swap, true)
		return err
	} else {
		return err
	}
}

// Utility function to perform bitonic sort in slice using concurrency
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

// Utility function to perform merge in bitonic sort using concurrency
func bitonicMergeC(l int, cnt int, slice reflect.Value, swap func(int, int), up bool) {
	if cnt > 1 {
		k := sizePartition(cnt)
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

// utility function returning greatest power of 2 less than n, used for partitioning the slices into powers of 2
func sizePartition(n int) int {
	k := 1
	for k > 0 && k < n {
		k = k << 1
	}
	return k >> 1
}
