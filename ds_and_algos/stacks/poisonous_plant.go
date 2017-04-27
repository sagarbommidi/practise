package main

import "fmt"

func max(a1, a2 int) int {
	if(a1 >= a2) {
		return a1
	} else {
		return a2
	}
}

func daysToBecomeImmortal(arr []int, size int) int {
	immortalArr := make([]int, size)
	immortalArr[0] = arr[0]
	index := 0
	aIndex := 0
	maxDays := 0
	hashMap := make(map[int]int)

	for x := 1; x < size; x++ {
		currentDays := hashMap[maxDays]
		if (arr[x] > arr[x-1]) {
			if (maxDays == 0) {
				maxDays = hashMap[maxDays]
			}
			hashMap[currentDays] = x-1
			continue
		}
		index++
		aIndex++

		if(x == 16) {
			fmt.Println(immortalArr)
		}
		for y := hashMap[maxDays]; y >= 0; y-- {
			fmt.Println(immortalArr[:hashMap[maxDays]])
			if(x == 16){
				fmt.Println(hashMap[maxDays])
				fmt.Println(arr[x], immortalArr[y], currentDays)
				//fmt.Println(arr[x], maxDays, currentDays)
			}
			currentDays++
			if (arr[x] > immortalArr[y]) {
				aIndex--
				hashMap[currentDays] = x
				maxDays = max(maxDays, currentDays)
				break
			}
		}
		immortalArr[index] = arr[x]
		fmt.Println(hashMap)
	}
	return maxDays
}

func main()  {
	fmt.Println("Implementing Poisonous Plants")
	//arrCount := 17
	var arr []int
	arr = []int {20, 5, 6, 15, 2, 2, 17, 2, 11, 5, 14, 5, 10, 9, 19, 12, 5} // Ans: 4
	//arr = []int {7, 6, 5, 4, 3, 2, 1}
	//arr = []int {6, 5, 8, 4, 7, 10, 9}
	fmt.Println(daysToBecomeImmortal(arr[:], len(arr)))
}