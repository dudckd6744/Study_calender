package main

import "fmt"

/*타입문 작성은 선택입니다*/

func calCoin(coinCount int, coin int) func() int { 
	return func()int{
		result := coinCount*coin
		return result
	}
}

func main() {
	var coin10, coin50, coin100, coin500 int
	fmt.Scan(&coin10, &coin50, &coin100, &coin500)
	
	
	
	add10 := calCoin(coin10, 10)()
	add50 := calCoin(coin50, 50)()
	add100 := calCoin(coin100, 100)()
	add500 := calCoin(coin500, 500)()
	
	totalmoney := add10+add50+add100+add500
	
	fmt.Println(totalmoney)	
}