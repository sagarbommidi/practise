package main

import (
	"fmt"
)

func max(a1, a2 int) int {
	if(a1 > a2) {
		return a1
	} else {
		return a2
	}
}

func swap(px, py *int) {
	tempx := *px
	tempy := *py
	*px = tempy
	*py = tempx
}

func quickSort(arr []int, left, right int) {
	if (left < right) {
		pivotIndex := partition(arr, left, right)
		quickSort(arr, left, pivotIndex - 1)
		quickSort(arr, pivotIndex + 1, right)
	}
}

func partition(arr []int, left, right int) int{
	pivot := arr[right]
	pindex := left
	for x := left; x < right; x++ {
		if(arr[x] <= pivot) {
			swap(&arr[pindex], &arr[x])
			pindex += 1
		}
	}
	swap(&arr[pindex], &arr[right])
	return pindex
}

func diff(a1, a2 int) int {
	if (a1 > a2) {
		return a1 - a2
	} else {
		return a2 - a1
	}
}

func getMinDiff(arr []int) int {
	fmt.Println(arr)
	minVal := diff(arr[1], arr[0])
	var val int
	for x := 2; x < len(arr); x++ {
		val = diff(arr[x], arr[x-1])
		fmt.Println(val)
		if(minVal == 0 || minVal > val) {
			minVal = val
		}
	}
	return minVal
}

func prefixSumMod(arr []int, mod int) int {
	size := len(arr)
	prefixArr := make([]int, size)
	prefixArr[0] = arr[0] % mod
	for x := 1; x < len(arr); x++ {
		prefixArr[x] = (prefixArr[x-1] + arr[x]) % mod
	}
	maxMod := 0
	for i := 0; i < size; i++ {
		for j := i-1; j > 0; j-- {
			maxMod = max(maxMod, ((prefixArr[i] - prefixArr[j] + mod) % mod))
		}
		maxMod = max(maxMod, prefixArr[i]);
	}
	return  maxMod
}

func main() {
	fmt.Println("Implementing Maximum subarray sum problem")
	arr := []int {3, 3, 9, 9, 5}
	mod := 7
	fmt.Println(prefixSumMod(arr, mod))
}

