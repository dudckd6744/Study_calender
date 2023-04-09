package main

import (
	"fmt"
)

func changeValue(num *int) {

	fmt.Println(num)
	*num = *num + 1
}

func main() {
	num := 1
	fmt.Println(num) // 출력: 1
	fmt.Println(&num)

	changeValue(&num)
	fmt.Println(num) // 출력: 2
}
