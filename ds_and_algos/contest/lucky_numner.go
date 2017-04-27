package main

import (
	"fmt"
)

func isLuckyNumber(val int) string {
	//if(val % 7 == 0 || val % 4 == 0 || (val % 7) % 4 == 0 || (val % 4) % 7 == 0) {
	////if(val % 7 == 0 || val % 4 == 0 || (val % 11) == 0) {
	//	return "YES"
	//} else {
	//	return "NO"
	//}
	if(val % 7 == 0 || val % 4 == 0 || val % 11 == 0 || val % 13 == 0) {
		return "Yes"
	} else {
		return "No"
	}

	return "No"
}

func main() {
	fmt.Println("ssss")
	fmt.Println(isLuckyNumber(26))
}

