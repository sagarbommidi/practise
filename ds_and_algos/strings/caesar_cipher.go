package main
import "fmt"

func encryptChar(char byte, key byte) string {
	newChar := char

	if(char >= 65 && char <= 90) {
		newChar += key
		if(newChar > 90) {
			diff := newChar - byte(90)
			newChar = byte(64) + diff
		}
	} else if(char >= 97 && char <= 122) {
		newChar += key
		if(newChar > 122) {
			diff := newChar - byte(122)
			newChar = byte(96) + diff
		}
	}
	return string(newChar)
}

func main() {
	fmt.Println("Implementation of caesar-cipher problem")
	str := "W"
	key := byte(87)
	key = key%26


	var charCount int
	_, _ = fmt.Scan(&charCount)

	var str string
	_, _ = fmt.Scan(&str)

	var key byte
	_, _ = fmt.Scan(&key)
	key = key%26

	output := ""
	for k:=0; k < len(str); k++ {
		output += encryptChar(str[k], key)
	}
	fmt.Println(output)
}
