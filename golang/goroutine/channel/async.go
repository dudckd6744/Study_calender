package main

import (
	"fmt"
	// "time"
)

func main() {
	chal := make(chan bool, 50)

	for i := 0; i < 20 ; i ++ {
		chal <- true
	}
	// close(chal)
	fmt.Println("송신 루틴 완료")
	
	for i := 0; i < 20 ; i ++ {
	  <- chal
	  fmt.Printf("수신한 데이터 %v\n", i)
	}
	// for i := range chal {
	// 	fmt.Println(i)
	// }
	
	// time.Sleep(time.Second * 3)
}