package main
import "fmt"

func main()  {
	var count int
	_, _ = fmt.Scan(&count)

	maxKey := -1
	var key, val int
	hashMap := make(map[int]int)
	for x := 1; x <= count; x++ {
		_, _ = fmt.Scan(&val)
		if(val == 0) {
			continue
		}
		key = x - val
		if(key <= 0) {
			key += count
			if (key == x) {
				continue
			}
			for j := x+1; j <= key; j++ {
				hashMap[j]++
			}
		} else {
			for i := 1; i <= key; i++ {
				hashMap[i]++
			}
			for i := x+1; i <= count; i++ {
				hashMap[i]++
			}
		}
		fmt.Println(hashMap[x])
	}
	//fmt.Println(hashMap)
	for key, val := range hashMap {
		if(maxKey == -1) {
			maxKey = key
		}
		if(val > hashMap[maxKey]) {
			maxKey = key
		} else if (val == hashMap[maxKey] && key < maxKey) {
			maxKey = key
		}
	}
	if(maxKey == -1) {
		maxKey = 1
	}
	fmt.Println(maxKey)
}
