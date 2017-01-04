package main

import (
	"fmt"
)

func mergeAndPrintStrings(str1, str2 string) {
	x := len(str1)
	y := len(str2)
	i := 0
	j := 0
	for ; (i < x) && (j < y); {
		if(str1[i] <= str2[j]) {
			fmt.Print(string(str1[i]))
			i++
		} else {
			fmt.Print(string(str2[j]))
			j++
		}
	}
	for ; i < x; i++ {
		fmt.Print(string(str1[i]))
	}
	for ; j < y; j++ {
		fmt.Print(string(str2[j]))
	}
	fmt.Print("\n")
}

func main() {
	//fmt.Println("Implementation of Morgan and a String")
	//str1 := "JACK"
	//str2 := "DANIEL"
	//output := mergeStrings(str1, str2, "")
	//fmt.Println(output == "DAJACKNIEL")
	//
	//output = mergeStrings("ABACABA", "ABACABA", "")
	//fmt.Println(output == "AABABACABACABA")

	var testCases int
	_, _ = fmt.Scan(&testCases)

	for x:=0; x<testCases; x++ {
		var str1 string
		var str2 string
		_, _ = fmt.Scan(&str1)
		_, _ = fmt.Scan(&str2)
		mergeAndPrintStrings(str1, str2)
	}
}