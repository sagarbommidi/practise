package main

import (
	"fmt"
	"strings"
	"strconv"
)

func min(d1, d2 int) int {
	if (d1 == 0) {
		return d2
	}else if(d1 >= d2) {
		return d2
	} else {
		return d1
	}
}

func getMinMaxinMap(hash map[int]int) []int {
	min := -1
	max := -1
	for key, _ := range hash {
		if(min == -1 || max == -1) {
			min = key
			max = key
		}
		if(key < min) {
			min = key
		}else if(key > max) {
			max = key
		}
	}
	return []int{min, max}
}

func main() {
	fmt.Println("Implementing AMazon maximize sales from techgig")
	hashMap := make(map[int]int)
	strArr := []string {
		"1 5 20",
		"3 8 15",
		"7 10 8",
	}
	for x:=0; x<len(strArr); x++ {
		arr := strings.Split(strArr[x], " ")
		end, _ := strconv.Atoi(arr[1])
		price, _ := strconv.Atoi(arr[2])
		for y, _ := strconv.Atoi(arr[0]); y <= end; y++ {
			hashMap[y] = min(hashMap[y], price)
		}
	}
	fmt.Println(hashMap)
	fmt.Println(getMinMaxinMap(hashMap))
	arr := getMinMaxinMap(hashMap)
	start := -1
	for x:=arr[0]; x<=arr[1]; x++ {
		if(start == -1) {
			start = x
		} else if (hashMap[x] != hashMap[start]) {
			fmt.Printf("%v %v %v\n", start, x-1, hashMap[start])
			start = x
		}
		if(x == arr[1]) {
			fmt.Printf("%v %v %v\n", start, x, hashMap[start])
		}
	}
}

