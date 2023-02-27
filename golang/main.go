package main

import (
	"fmt"
	"time"
)

func main() {
	ch:= make(chan string)

	go sendMessage(ch)

	for i := 10; i > 0; i--{
		select {
			case msg:= <- ch :
				fmt.Printf("%s 메세지가 발송되었습니다.", msg)
				return
			default :
				fmt.Println(i,"초안에 메시지를 입력하세요.")
		}
		time.Sleep(1 * time.Second)
	}
}

func sendMessage(ch chan string) {
	var input string;

   	fmt.Scanln(&input)

	ch <- input

	close(ch)
}