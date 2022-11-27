package main

import "fmt"

// defer 연속적으로 출력시 lifo 방식으로 출력됨
func main() {
	
	var name string
	var names []string
	
	for {
		fmt.Scanln(&name)
		if name == "0"{
			break
		} else {
			names = append(names, name)
		}
	}
	
	for _, v := range names {
		defer fmt.Println(v)
	}
	
}