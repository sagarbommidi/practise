package main

import (
	"fmt"
	"strings"
	"strconv"
)

func MakeRich(intArr []int, size, k int) ([]int, string) {
	index := size-1
	var times, extra int
	indArr := make([]int, size/2+1)
	for i := 0; i < size/2; i++ {
		if(intArr[i] != intArr[index-i]) {
			indArr[times] = i
			times++
		}
	}
	if(times > k) {
		return nil, "-1"
	}
	extra = k - times
	for j := 0; j < times; j++ {
		in := indArr[j]
		if(extra > 0 && intArr[in] != 9 && intArr[index-in] != 9) {
			intArr[indArr[j]] = 9
			extra--
		}
		if(intArr[in] > intArr[size-1-in]) {
			intArr[index - in] = intArr[in]
		} else {
			intArr[in] = intArr[index-in]
		}
	}
	i := 0
	for extra > 0 {
		if(extra == 1 && size % 2 == 1) {
			extra--
			intArr[size/2] = 9
			break
		}
		if(extra % 2 == 0) {
			for ; i < size/2; i++ {
				if(intArr[i] != 9) {
					intArr[i] = 9
					intArr[index-i] = 9
					extra -= 2
					break
				}
			}
		}
	}

	return intArr, ""
}

func main() {
	var size, k int
	var str string

	_, _ = fmt.Scan(&size, &k)
	_, _ = fmt.Scan(&str)

	strArr := strings.Split(str, "")
	intArr := make([]int, size)

	for x := 0; x < size; x++ {
		intArr[x], _ = strconv.Atoi(strArr[x])
	}

	intArr, err := MakeRich(intArr, size, k)
	if (err != "") {
		fmt.Println(err)
	} else {
		for x := 0; x < size; x++ {
			fmt.Print(intArr[x])
		}
	}
}
