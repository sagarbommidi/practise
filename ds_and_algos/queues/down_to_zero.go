package main

import (
	"fmt"
	"math"
)

func max(n1, n2 int) int {
	if(n1 > n2) {
		return n1
	} else {
		return n2
	}
}

func TransformedNumber(num int) int {
	sRoot := rune(math.Sqrt(float64(num)))
	for x := sRoot; x > 1; x-- {
		if (num % x == 0) {
			return max(x, num/x)
		}
	}
	return num-1
}

func DownToZero(num int) int {
	if (num == 0) {
		return 0
	}
	if (num == 1) {
		return 1
	}
	if (num == 2) {
		return 2
	}
	return 2 + DownToZero(TransformedNumber(num))
}

func main() {
	fmt.Println("Implementing Down to Zero - II in queues")
	var count, val int
	_, _ = fmt.Scan(&count)

	for x := 0; x < count; x++ {
		_, _ = fmt.Scan(&val)
		fmt.Println(DownToZero(val))
	}
}
