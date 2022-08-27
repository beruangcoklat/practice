package main

func _mergeSort(arr []int, l, r int) {
	if l == r {
		return
	}

	m := (l + r) / 2
	_mergeSort(arr, l, m)
	_mergeSort(arr, m+1, r)

	leftArr := []int{}
	rightArr := []int{}

	for i := l; i <= m; i++ {
		leftArr = append(leftArr, arr[i])
	}

	for i := m + 1; i <= r; i++ {
		rightArr = append(rightArr, arr[i])
	}

	idxL, idxR := 0, 0
	idx := l
	for idxL < len(leftArr) && idxR < len(rightArr) {
		min := 0
		if leftArr[idxL] < rightArr[idxR] {
			min = leftArr[idxL]
			idxL++
		} else {
			min = rightArr[idxR]
			idxR++
		}
		arr[idx] = min
		idx++
	}

	idxLast := idxR
	arrLast := rightArr
	if idxL < len(leftArr) {
		idxLast = idxL
		arrLast = leftArr
	}

	for idxLast < len(arrLast) {
		arr[idx] = arrLast[idxLast]
		idxLast++
		idx++
	}
}

func mergeSort(arr []int) {
	_mergeSort(arr, 0, len(arr)-1)
}
