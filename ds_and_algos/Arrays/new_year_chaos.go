package main

import "fmt"

func isChaotic(arr []int, size int) int {
	swapCount := 0
	var temp int
	for x := 0; x < size; x++ {
		if( (arr[x] < (x-1)) || arr[x] > (x+3)) {
			return -1
		} else if(arr[x] == x+1) {
			continue
		} else if(arr[x] > x+1) {
			for y := x; y < arr[x]; y++ {
				fmt.Println("y: ", y)
				fmt.Println(arr)
				temp = arr[y]
				arr[y] = arr[y+1]
				arr[y+1] = temp
				swapCount++
			}
		} else {
			continue
		}
	}
	return swapCount
}
func main()  {
	fmt.Println("Implementing New year chaos")
	var arr []int
	//arr = []int {2, 1, 5, 3, 4} // Ans: 3
	arr = []int {2, 5, 1, 3, 4} // Ans: Too chaotic
	res := isChaotic(arr, len(arr))
	if(res == -1) {
		fmt.Println("Too chaotic")
	} else {
		fmt.Println(res)
	}
}
