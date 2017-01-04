package main

import "fmt"

func max(a1, a2 int) int {
	if(a1 > a2) {
		return a1
	} else {
		return a2
	}
}

func min(a1, a2 int) int {
	if (a1 == 0 || a2 < a1) {
		return a2
	} else {
		return a1
	}
}

func diff(a1, a2 int) int {
	if(a1 > a2) {
		return a1 - a2
	} else {
		return a2 - a1
	}
}

func isValid(str string) string {
	strMap := make(map[string]int)
	for x := 0; x < len(str); x++ {
		strMap[string(str[x])] += 1
	}
	countMap := make(map[int]int)
	for _, val := range strMap {
		countMap[val]++
	}
	countArr := make([]int, len(countMap))
	i := 0
	for key, _ := range countMap {
		countArr[i] = key
		i++
	}

	if(len(countArr) == 1) {
		return "YES"
	} else if(len(countArr) > 2) {
		return "NO"
	} else if(countMap[1] > 1) {
		return "NO"
	} else if(countMap[1] == 1) {
		return "YES"
	} else if(diff(countArr[0], countArr[1]) == 1 && (countMap[countArr[0]] == 1 || countMap[countArr[1]]== 1)) {
		return "YES"
	} else {
		return "NO"
	}
}
func main() {
	fmt.Println("Implementing Sherlock and valid string program")
	//str := "hfchdkkbfifgbgebfaahijchgeeeiagkadjfcbekbdaifchkjfejckbiiihegacfbchdihkgbkbddgaefhkdgccjejjaajgijdkd"
	//ANS: YES

	//str := "ibfdgaeadiaefgbhbdghhhbgdfgeiccbiehhfcggchgghadhdhagfbahhddgghbdehidbibaeaagaeeigffcebfbaieggabcfbiiedcabfihchdfabifahcbhagccbdfifhghcadfiadeeaheeddddiecaicbgigccageicehfdhdgafaddhffadigfhhcaedcedecafeacbdacgfgfeeibgaiffdehigebhhehiaahfidibccdcdagifgaihacihadecgifihbebffebdfbchbgigeccahgihbcbcaggebaaafgfedbfgagfediddghdgbgehhhifhgcedechahidcbchebheihaadbbbiaiccededchdagfhccfdefigfibifabeiaccghcegfbcghaefifbachebaacbhbfgfddeceababbacgffbagidebeadfihaefefegbghgddbbgddeehgfbhafbccidebgehifafgbghafacgfdccgifdcbbbidfifhdaibgigebigaedeaaiadegfefbhacgddhchgcbgcaeaieiegiffchbgbebgbehbbfcebciiagacaiechdigbgbghefcahgbhfibhedaeeiffebdiabcifgccdefabccdghehfibfiifdaicfedagahhdcbhbicdgibgcedieihcichadgchgbdcdagaihebbabhibcihicadgadfcihdheefbhffiageddhgahaidfdhhdbgciiaciegchiiebfbcbhaeagccfhbfhaddagnfieihghfbaggiffbbfbecgaiiidccdceadbbdfgigibgcgchafccdchgifdeieicbaididhfcfdedbhaadedfageigfdehgcdaecaebebebfcieaecfagfdieaefdiedbcadchabhebgehiidfcgahcdhcdhgchhiiheffiifeegcfdgbdeffhgeghdfhbfbifgidcafbfcd"
	// ANS: YES

	//str := "aabbc"

	str := "hfchdkkbfifgbgebfaahijchgeeeiagkadjfcbekbdaifchkjfejckbiiihegacfbchdihkgbkbddgaefhkdgccjejjaajgijdkd"
	fmt.Println(isValid(str))
}
