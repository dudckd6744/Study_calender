package main

import "fmt"

func main() {
	var a, b = 10, 5
	var result int
	
	c := make(chan int)
	
	go func() {
		c <- a + b
	}()
	
	result = <-c
	fmt.Printf("두 수의 합은 %d입니다.", result)
}