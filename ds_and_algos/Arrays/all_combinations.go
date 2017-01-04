package main

import "fmt"

func getCombinations(leftArr, rightArr []int, resultArr [][]int) [][]int {
	if(len(rightArr) >= 1) {
		for x := 0; x < len(rightArr); x++ {
			temp := append(leftArr, rightArr[x])
			resultArr = append(resultArr, temp)
			if (x + 1 < len(rightArr)) {
				getCombinations(temp, rightArr[x + 1:], resultArr)
			}
		}
	}
	return resultArr
}

func allCombinationsOfArr(arr []int) {
	size := len(arr)
	resultArr := make([][]int, 0)
	newArr := make([]int, 0, size)
	fmt.Println(getCombinations(newArr, arr, resultArr))
}

func main() {
	fmt.Println("Implementing Maximum subarray sum problem")
	//arr := []int {3, 3, 9, 9, 5}
	//mod := 7
	//fmt.Println(findMaxSumArray(arr, mod))
	arr := []int {1,2,3,4,5}
	allCombinationsOfArr(arr)
}
