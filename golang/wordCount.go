package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I am learning Go!"
	strArray := strings.Split(str, " ")
	var mapArr = make(map[string]int)
	for _, v := range strArray {
		mapArr[v] += 1
	}
	for key, value := range mapArr {
		fmt.Println(key, " ==> ", value)
	}
}
