package main

import (
	. "fmt"
)

// 입력한 과목 이름과 점수가 각 줄에 출력됩니다. 순서는 입출력 예시와 다르게 섞여서 출력될 수 있습니다.
// 마지막 줄에는 점수들의 평균이 소수점 두 번째 자리까지 출력됩니다.


func main_02() {
	var m = make(map[string]int)

	var key string
	var value int
	
	for {
		Scan(&key)
		if key == Sprint(0) {
			goto ONE
		}
		Scan(&value)
		m[key] = value
	}

	ONE:
	var sum int
	var result float64
	for key, v := range m {
		sum += v
		Println(key , v)
	}


	result = float64(sum)/float64(len(m))

	Printf("%0.2f",result)

}