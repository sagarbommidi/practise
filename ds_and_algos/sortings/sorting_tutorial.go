package main

import (
	"fmt"
)

func searchInSortedArr(arr []int, key, start, end int) int {
	mid := (start + end-1)/2
	if(arr[mid] == key) {
		return mid
	}
	if((start == end) && arr[mid] != key) {
		return -1
	}
	if (key <= arr[mid]) {
		return searchInSortedArr(arr, key, 0, mid)
	} else {
		return searchInSortedArr(arr, key, mid+1, end)
	}
}

func main() {
	var searchKey int
	_, _ = fmt.Scan(&searchKey)

	var arrCount int
	_, _ = fmt.Scan(&arrCount)

	arr := make([]int, arrCount)
	for x:=0; x<arrCount; x++ {
		var val int
		_, _ = fmt.Scan(&val)
		arr[x] = val
	}

	fmt.Println(searchInSortedArr(arr, searchKey, 0, arrCount))
}