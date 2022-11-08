package main

import (
	. "fmt"
)

func main() {
	names := make([]string, 0 , 10)
	var name string
	
	// 입력한 이름 중에 가장 긴 이름과 그 이름의 길이가 띄어쓰기돼서 출력됩니다.
	for {
		Scan(&name)
		if name == Sprint(1) { 
			goto ONE
		}
		names = append(names, name)
	}
	ONE:
	var result string = names[0]
		for _, v := range names {
			if len(result) < len(v) {
				result = v
			}
		}
		
	Println(result, len(result) )
}