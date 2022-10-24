package test

import "fmt"

var Art = "test"
// 두수의 차
// 무조건 큰 수에서 작은 수를 뺀 결과를 출력해야합니다.
func Test()int {
	
	var num1,num2,result int

	fmt.Scan(&num1,&num2)

	if num1 > num2 {
		result = num1 - num2
	} else {
		result = num2 - num1
	}
	
	fmt.Print(result)
	return result
}