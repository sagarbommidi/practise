package main

import "fmt"

func AltCharsString(str string) int {
	prevChar := str[0]
	delCount := 0
	for x := 1; x < len(str); x++ {
		if(str[x] == prevChar) {
			delCount++
		} else {
			prevChar = str[x]
		}
	}
	return delCount
}

func main() {
	fmt.Println("Implementing Alternating Cahracters")
	var testCases int
	_, _ = fmt.Scan(&testCases)

	for x:=0; x<testCases; x++ {
		var str string
		_, _ = fmt.Scan(&str)
		fmt.Println(AltCharsString(str))
	}
}
