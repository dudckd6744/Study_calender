package main

import "fmt"

func main() {
	
	var multiArray [2][3][4]int //3차원 배열 선언
    multiArray[1][1][2] = 10  // 인덱스를 이용한 값 초기화

	fmt.Println((multiArray))


}
[
	[
		[0 0 0 0] [0 0 0 0] [0 0 0 0]
	] 
	[
		[0 0 0 0] [0 0 10 0] [0 0 0 0]
	]
]