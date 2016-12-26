package main
import "fmt"

func strToMap(str string) map[byte]int {
	strMap := make(map[byte]int)
	size := len(str)
	for x := 0; x < size; x++ {
		strMap[str[x]] += 1
	}
	return strMap

}
func diffOfNums(n1, n2 int) int {
	if(n1 >= n2) {
		return (n1 - n2)
	} else {
		return (n2 - n1)
	}
}
func changeCountToMakeAnagram(str string) int {
	changeCount := 0
	size := len(str)
	if (size & 1 == 1) {
		return -1
	}
	str1Map := strToMap(str[:(size/2)])
	str2Map := strToMap(str[(size/2):])
	var diff int
	for key, val := range str2Map {
		diff = diffOfNums(val, str1Map[key])
		changeCount += diff
	}
	return changeCount
}

func main() {
	fmt.Println("Implementing Anagram functionality")
	var testCases int
	_, _ = fmt.Scan(&testCases)

	for x:=0; x<testCases; x++ {
		var str string
		_, _ = fmt.Scan(&str)
		fmt.Println(changeCountToMakeAnagram(str))
	}
}
