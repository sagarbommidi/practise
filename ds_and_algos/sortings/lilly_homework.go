package main

import "fmt"

func selectionSort(arr []int) int {
	//def selection_sort(arr)
	//for i in (0...arr.length) do
	//min = i
	//for j in (i+1...arr.length) do
	//min = j if arr[j] < arr[min]
	//end
	//arr[i], arr[min] = arr[min], arr[i]
	//end
	//end
	//
	size := len(arr)
	swaps := 0
	for x := 0; x < size; x++ {
		min := x
		for y := x+1; y < size; y++ {
			if(arr[y] < arr[x]) {
				min = y
			}
		}
		if(min != x) {
			temp := arr[x]
			arr[x] = arr[min]
			arr[min] = temp
			swaps++
		}
	}
	return swaps
}

func main() {
	fmt.Println("Lilly homework problem")
	arr := []int {2, 5, 3, 1}
	fmt.Println(selectionSort(arr))
}
