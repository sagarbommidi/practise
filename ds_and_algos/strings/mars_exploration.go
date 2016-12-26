package main

import (
	"fmt"
)

func messageTrasmissionCount(str string) int{
	alteredCount := 0
	size := len(str)
	testStr := "SOS"
	for x:=0; x<size; x++ {
		if(str[x] != testStr[x%3]) {
			alteredCount++
		}
	}
	return alteredCount
}

func main() {
	fmt.Println("Implementing Mars Exploration")
	var str string
	_, _ = fmt.Scan(&str)
	fmt.Println(messageTrasmissionCount(str))
}

