package main

import "fmt"

func countWords(str string) int{
	count := 1
	for x := 0; x <len(str); x++ {
		if(str[x] >= 65 && str[x] <= 90 ) {
			count += 1
		}
	}
	return count
}

func main() {
	var input string
	_, _ = fmt.Scan(&input)
	//input := "saveChangesInTheEditor"
	fmt.Println(countWords(input))
}
