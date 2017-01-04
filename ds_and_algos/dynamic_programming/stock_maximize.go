package main

import "fmt"

type Market struct {
	Invested int
	Returns int
	SharesCount int
}

func (m *Market)BuyOneShare(price int) {
	m.SharesCount += 1
	m.Invested += price
}

func (m *Market)getProfit() int {
	return (m.Returns - m.Invested)
}

func (m *Market) SellShares(price int) {
	m.Returns += m.SharesCount * price
	m.SharesCount = 0
}

func getMax(a1, a2 int) int {
	if(a1 >= a2) {
		return a1
	} else {
		return a2
	}
}

func dpMaxElementArray(arr []int, size int) []int {
	maxArr := make([]int, size)
	maxArr[size-1] = arr[size-1]
	for k:=size-2; k>=0; k-- {
		maxArr[k] = getMax(maxArr[k+1], arr[k])
	}
	return maxArr
}

func predictStrategy(arrEle, maxArrEle, size int) string {
	strategy := "nothing"
	if(arrEle == maxArrEle) {
		strategy = "sell"
	} else if (arrEle < maxArrEle) {
		strategy = "buy"
	}
	return strategy
}

func (market *Market)executeSales(priceList []int, count int) {
	maxArr := dpMaxElementArray(priceList, count)
	for x := 0; x < count; x++ {
		switch predictStrategy(priceList[x], maxArr[x], count) {
		case "buy":
			market.BuyOneShare(priceList[x])
		case "sell":
			market.SellShares(priceList[x])
		}
	}

}

func main() {
	fmt.Println("Implementing Sales Strategy for maximizing the profits")
	market := Market{}
	prices := []int {1, 2, 100}
	market.executeSales(prices, len(prices))
	fmt.Println(market.getProfit())

	market = Market{}
	prices = []int {5, 3, 2}
	market.executeSales(prices, len(prices))
	fmt.Println(market.getProfit())

	market = Market{}
	prices = []int {1, 3, 1, 2}
	market.executeSales(prices, len(prices))
	fmt.Println(market.getProfit())

	//var testCases int
	//_, _ = fmt.Scan(&testCases)
	//
	//for x:=0; x<testCases; x++ {
	//	var priceCount int
	//	_, _ = fmt.Scan(&priceCount)
	//
	//	arr := make([]int, priceCount)
	//	for y := 0; y < priceCount; y++ {
	//		var val int
	//		_, _ = fmt.Scan(&val)
	//		arr[y] = val
	//	}
	//	market := Market{}
	//	market.executeSales(arr, priceCount)
	//	fmt.Println(market.getProfit())
	//}

}