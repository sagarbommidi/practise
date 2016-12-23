package main

import (
	"fmt"
)

func reverse(str string) string {
	output := ""
	for x := len(str)-1; x >= 0; x-- {
		output += string(str[x])
	}
	return output
}

func getDiff(char1, char2 byte) byte {
	if(char1 > char2) {
		return char1 - char2
	} else {
		return char2 - char1
	}
}

func checkFunnyString(str string) string {
	revStr := reverse(str)
	funnyString := "Funny"
	for x:=0; x<len(str)-1; x++ {
		if(getDiff(str[x], str[x+1]) != getDiff(revStr[x], revStr[x+1])) {
			funnyString = "Not Funny"
			break
		}
	}
	return funnyString
}

func main() {
	var testCases int
	_, _ = fmt.Scan(&testCases)

	for x:=0; x<testCases; x++ {
		var str string
		_, _ = fmt.Scan(&str)
		fmt.Println(checkFunnyString(str))
	}

}
