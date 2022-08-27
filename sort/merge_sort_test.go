package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	t.Parallel()

	rand.Seed(time.Now().UnixNano())
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}

	mergeSort(arr)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			t.Error("merge sort failed")
			break
		}
	}
}
