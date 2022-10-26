package main

import "fmt"

func main() {
	
	num1:=2

	for i :=1; i <10; i++ {
		if num1 % 2 ==1 {
			for k:=1; k <=num1; k++ {
				result := num1 * k
				fmt.Printf("%d x %d = %d\n",num1, k, result)
			}
			fmt.Println("")

		}
		num1++

	}
}