package main

import "fmt"

func main_01() {
	
	var result int
	
	for i := 0 ; i < 10 ; i++ {
		for j := 9 ; j >= 0  ; j-- {
			result = 99

			if i == j {
			continue
			}
			if i+j == 9 {
			fmt.Printf("%d%d + %d%d = %d\n", i, j, j, i, result)
			}
		}	
	}
}