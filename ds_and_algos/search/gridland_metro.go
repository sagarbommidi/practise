package main

import (
	"fmt"
)

func diff(a1, a2 int) int {
	if (a1 == a2) {
		return 0
	} else if (a1 > a2) {
		return a1 - a2
	} else {
		return a2 - a1
	}
}

func getTrackLength(hash map[int]int) int {
	fmt.Println("hash: ", hash)
	var keys []int
	for k := range hash {
		keys = append(keys, k)
	}
	fmt.Println("keys: ", keys)
	sum := hash[keys[0]]
	for x := 1; x < len(keys); x++ {
		startx := keys[x]
		endx := startx + hash[keys[x]] - 1
		y := x-1
		val := 0
		for ; y >= 0; {
			starty := keys[y]
			endy := starty + hash[keys[y]] - 1
			if(starty <= startx && endy >= endx) {
				val += (diff(startx, starty) + diff(endx, endy))
			} else if( starty >= startx && starty <= endx && endy > endx ) {
				val += diff(endx, endy)
			} else if (starty >= startx && endy <= endx) {
				continue
			} else {
				val += hash[y]
			}
		}
		sum += val
	}
	return sum
}

func main() {
	var row, col, tracks int
	_, _ = fmt.Scan(&row, &col, &tracks)

	tracksMap := make(map[int]map[int]int)

	for x := 0; x < tracks; x++ {
		var rowNum, col1, col2 int
		_, _ = fmt.Scan(&rowNum, &col1, &col2)
		val, ok := tracksMap[rowNum]
		if !ok {
			val = make(map[int]int)
			tracksMap[rowNum] = val
		}
		tracksMap[rowNum][col1] = (col2 - col1 + 1)
	}
	occupied := 0
	for _, val := range tracksMap {
		occupied += getTrackLength(val)
	}
	res := (row * col) - occupied
	fmt.Println(res)
}
