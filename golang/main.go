package main

import "fmt"

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