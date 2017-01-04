package main

import "fmt"

func minIndex(arr []int, size int) int {
	min := 0
	for k:=1; k < size; k++ {
		if(arr[k] < arr[min]) {
			min = k
		}
	}
	return min
}

func maxIndex(arr []int, size int) int {
	max := 0
	for k:=1; k < size; k++ {
		if(arr[k] > arr[max]) {
			max = k
		}
	}
	return max
}

func increaseArrayElements(arr []int, size, val, maxIndex int) []int {
	for k:=0; k<size; k++ {
		if(k == maxIndex) {
			continue
		}else {
			arr[k] += val
		}
	}
	return arr
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

func makeEqual(min, max, count, increaseBy int) (int, int, int) {
	fmt.Println(max, "-", min, "=", max-min)
	for max-min > 0 {
		diff := max - min
		switch {
		case diff >= 5:
			count += diff / 5
			increaseBy += diff / 5 * 5
			min += diff / 5 * 5
		case diff >= 2:
			count += diff / 2
			increaseBy += diff / 2 * 2
			min += diff / 2 * 2
		case diff == 1:
			count += 1
			increaseBy += 1
			min += 1
		case diff == 0:
			break
		}
	}
	return min, count, increaseBy
}

func equalChacos(arr []int, size int) int{
	count := 0
	increaseBy := 0
	quickSort(arr, 0, size-1)

	lastMin := arr[0]
	for x := 1; x < size; x++ {
		fmt.Println(lastMin, increaseBy)
		lastMin, count, increaseBy = makeEqual(lastMin, arr[x]+increaseBy, count, increaseBy)
	}
	return count
}

func main() {
	fmt.Println("Implementing equal chacolates in Dynamic Programming in HackerRank")
	//var testCases int
	//_, _ = fmt.Scan(&testCases)
	//
	//for k := 0; k < testCases; k++ {
	//	var internCount int
	//	_, _ = fmt.Scan(&internCount)
	//
	//	arr := make([]int, internCount)
	//	for y := 0; y < internCount; y++ {
	//		var val int
	//		_, _ = fmt.Scan(&val)
	//		arr[y] = val
	//	}
	//	fmt.Println(equalChacos(arr, internCount))
	//}
	arr := []int {1, 5, 5, 10, 10}
	fmt.Println(equalChacos(arr, len(arr)))
}
