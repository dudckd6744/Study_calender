package main

import "fmt"

func main() {

	code := "abc1abc1abc"

	mode := 0
	let := ""
	for i, v := range code {
		if mode == 0 {
			if string(v) != "1" {
				if i%2 == 0 {
					let += string(code[i])
				}
			} else {
				mode = 1
			}
		} else {
			if string(v) != "1" {
				if i%2 == 1 {
					let += string(code[i])
				}
			} else {
				mode = 0
			}
		}
	}

	if let == "" {
		fmt.Print("EMPTY")
	} else {
		fmt.Print(let)
	}
}
