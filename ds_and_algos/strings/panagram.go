package main
import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')

	strHash := make(map[string]int)
	for x := 0; x < len(str); x++ {
		char := strings.ToLower(string(str[x]))
		if(char != " ") {
			strHash[char] += 1
		}
	}

	if(len(strHash) == 26) {
		fmt.Println("pangram")
	} else {
		fmt.Println("not pangram")
	}
}
