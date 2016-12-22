// https://www.hackerrank.com/challenges/reduced-string

package main

import (
	"fmt"
	"strings"
)

func reduceString(arr []string, output string) string {
	if(len(arr) == 0) {
		return output
	}
	if(len(arr) == 1){
		output+= arr[0]
		return output
	}
	if(arr[0] == arr[1]) {
		return reduceString(arr[2:], output)
	}else {
		output += arr[0]
		return reduceString(arr[1:], output)
	}
	return output
}

func haveRepeatedChars(arr []string) bool {
	for x:=0; x < len(arr)-1; x++ {
		if(arr[x] == arr[x+1]) {
			return true
		}
	}
	return false
}

func main(){
	var input string
	_, _ = fmt.Scan(&input)
	output := ""
	output = reduceString(strings.Split(input, ""), "")
	for haveRepeatedChars(strings.Split(output, "")) {
		output = reduceString(strings.Split(output, ""), "")
	}
	if(len(output) == 0) {
		fmt.Println("Empty String")
	} else {
		fmt.Println(output)
	}
}
