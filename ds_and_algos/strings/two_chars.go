package main

import (
	"fmt"
)

func twoCharString(str string) string {
	max_output := ""
	output := ""
	uniqChars := getUniqChars(str)
	for x:=0; x < len(uniqChars)-1; x++ {
		for y:=x+1; y<len(uniqChars); y++ {
			key := []byte{uniqChars[x], uniqChars[y]}
			output = string(key[0])
			i := 0
			count := 0
			for ; i<len(str); i++ {
				if(str[i] == key[0]) {
					count++
				}else if(str[i] == key[1]) {
					output += string(key[1])
					i++
					break
				}
			}
			if(count > 1) {
				continue
			}
			for ; i < len(str); i++ {
				if ((str[i] != output[len(output)-1]) && (str[i] != output[len(output)-2])) {
					continue
				} else if (str[i] == output[len(output)-1]) {
					output = ""
					break
				} else if (str[i] == output[len(output)-2]) {
					output += string(str[i])
					continue
				}
			}
			if (len(max_output) < len(output)) {
				max_output = output
			}
		}
	}
	return max_output
}

func getUniqChars(str string) string {
	output := ""
	strMap := make(map[string]int)
	for x:=0; x < len(str); x++ {
		if(strMap[string(str[x])] == 0) {
			strMap[string(str[x])] += 1
			output += string(str[x])
		}
	}
	return output
}

func main() {
	//var input string
	//_, _ = fmt.Scan(&input)
	input := "muqqzbcjmyknwlmlcfqjujabwtekovkwsfjrwmswqfurtpahkdyqdttizqbkrsmfpxchbjrbvcunogcvragjxivasdykamtkinxpgasmwz"
	fmt.Println(len(twoCharString(input)))
}
