package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string
	var result string
	fmt.Scan(&s1)

	num := len(s1)

	for i := 0; i < num; i++ {
		if strings.ToUpper(s1[i:i+1]) == s1[i:i+1] {
			result += strings.ToLower(s1[i : i+1])
		}
		if strings.ToLower(s1[i:i+1]) == s1[i:i+1] {
			result += strings.ToUpper(s1[i : i+1])
		}

	}
	fmt.Print((result))

}
