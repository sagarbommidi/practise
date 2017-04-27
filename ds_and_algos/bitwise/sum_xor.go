package main

import "fmt"

func main() {
	fmt.Println("Implementing Sum vs XOR")
	//var val int
	//_, _ = fmt.Scan(&val)

	val := 1000000000000000
	count := 0
	for x := 0; x <= val; x++ {
		if((val + x) == (val ^ x)) {
			fmt.Println((val + x), (val ^ x))
			count++
		}
	}
	fmt.Println(count)
}
