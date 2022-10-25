package main

import (
	"fmt"
)

func main() {
	
	var input  int
	var result,num1,num2 float32

	fmt.Scan(&input)
	fmt.Scan(&num1,&num2)

	switch input {
	case 1:
		result = float32(num1+num2)
	case 2:
		result = float32(num1-num2)
	case 3:
		result = float32(num1*num2)
	case 4:
		result = float32(num1/num2)
	default :
		fmt.Print("잘못된입력입니다")
		return
	}
	
	fmt.Printf("%.1f\n", result)
}