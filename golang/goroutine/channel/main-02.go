package main

import "fmt"
 
func add(a int, b int ,ch chan int) {
	ch <- a + b
}
//채널에 개념 잡기
func main() {
	var num1, num2 int	
	
	ch := make(chan int)
	
	fmt.Scanln(&num1, &num2)
	 
	
 	go add(num1,num2 ,ch)
	
	fmt.Println(<-ch)
}
