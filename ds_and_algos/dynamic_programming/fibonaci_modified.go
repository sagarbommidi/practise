package main

import (
	"fmt"
	"strconv"
)

func max(a1, a2 int) int {
	if (a1 >= a2) {
		return a1
	} else {
		return a2
	}
}

func getIntegerArray(str string) []int {
	size := len(str)
	arr := make([]int, size)
	for x := 0; x < size; x++ {
		arr[x], _ = strconv.Atoi(string(str[x]))
	}
	return arr
}

func arrToString(arr []int, size int) string {
	str := ""
	intFound := 0
	for i := 0; i < size; i++ {
		if (intFound == 0 && arr[i] != 0) {
			intFound = 1
		}
		if (intFound == 1) {
			str += strconv.Itoa(arr[i])
		}
	}
	return str
}

func SumOfTwoIntegerStrings(str1, str2 string) string {
	strArr1 := getIntegerArray(str1)
	strArr2 := getIntegerArray(str2)
	size1 := len(strArr1)
	size2 := len(strArr2)
	size := max(size1, size2) + 1
	sumArr := make([]int, size)
	x := size1 - 1
	y := size2 - 1
	z := size - 1
	carry := 0
	sum := 0
	for ; z >= 0 ; z-- {
		sum = carry
		if(x >= 0) {
			sum += strArr1[x]
			x--
		}
		if(y >= 0) {
			sum += strArr2[y]
			y--
		}
		carry = (sum) / 10
		sumArr[z] = sum % 10
	}
	return arrToString(sumArr, size)
}

func multOfTwoIntegerStrings(str1, str2 string) string {
	strArr1 := getIntegerArray(str1)
	strArr2 := getIntegerArray(str2)
	size1 := len(strArr1)
	size2 := len(strArr2)
	size := (max(size1, size2) + 1) * 2
	multiArr := make([]int, size)

	var z int
	mult := 0
	for x := size1-1; x >= 0; x-- {
		z = size + x - size1
		for y := size2-1; y >= 0; y-- {
			mult = multiArr[z] + (strArr1[x] * strArr2[y])
			multiArr[z] = mult % 10
			z--
			multiArr[z] += mult /10
		}
	}
	return arrToString(multiArr, size)
}

func dpFibonaci(t1, t2 string, n int) string {
	arrMap := make([]string, n)
	arrMap[0] = t1
	arrMap[1] = t2

	var square string
	for k := 2; k < n; k++ {
		square = multOfTwoIntegerStrings(arrMap[k-1], arrMap[k-1])
		arrMap[k] = SumOfTwoIntegerStrings(square, arrMap[k-2])
	}
	return arrMap[n-1]
}
func main() {
	n := dpFibonaci("2", "2", 20)
	fmt.Printf(n)
	//var t1, t2 string
	//var n int
	//_, _ = fmt.Scan(&t1, &t2, &n)
	//fmt.Println(dpFibonaci(t1, t2, n))
}