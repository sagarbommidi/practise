package main

import "fmt"

func isAnagramAPalindrom(str string) string {
	stringMap := make(map[byte]int)
	size := len(str)
	for x := 0; x < size; x++ {
		stringMap[str[x]] += 1
	}
	oddChars := 0
	for _, val := range stringMap {
		if(val & 1 == 1) {
			oddChars++
		}
	}
	if((size & 1 == 0) && oddChars == 0) || ((size & 1 == 1) && oddChars == 1) {
		return "YES"
	} else {
		return "NO"
	}

}

func main() {
	fmt.Println("Implementing Anagram is a polindrom or not")
	var str string
	_, _ = fmt.Scan(&str)
	fmt.Println(isAnagramAPalindrom(str))
}
