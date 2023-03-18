package main

import (
	"fmt"
	"sync"
	// "time"
)

func landom(i int,wait *sync.WaitGroup){
	// defer wait.Done()
	fmt.Printf("%v \n",i)
}

func main(){
	wait := new(sync.WaitGroup)

	wait.Add(10)

	for i := 0; i < 10; i++ {

		go landom(i,wait)
	}
	// wait.Wait()
}

// type item struct {
// 	name string
// 	price int
// 	amount int
// }

// type buyer struct {
// 	point int
// 	shippingBucket map[string]int
// }

// func newBuyer() *buyer {
// 	d:= buyer{}
// 	d.point = 1000000
// 	d.shippingBucket = map[string]int{}

// 	return &d
// }

// func buyItem (choice int, items []item, buyer *buyer){

// 	fmt.Print("수량을 입력해주세요")

// 	var buyCount int

// 	fmt.Scan(buyCount)

//     choiceItem := items[choice] 

// 	totalPrice := choiceItem.price * buyCount

// 	items[choice].price -= totalPrice	
// 	items[choice].amount -= buyCount

// }

// func main() {

// 	items := make([]item, 5)
// 	buyer := newBuyer()

// 	items[0] = item{"텀블러", 10000, 30}
// 	items[1] = item{"롱패딩", 500000, 20}
// 	items[2] = item{"투미 백팩", 400000, 20}
// 	items[3] = item{"나이키 운동화", 150000, 50}
// 	items[4] = item{"빼빼로", 1200, 500}

// 	for {
// 		menu := 0 // 첫 메뉴
	
// 		fmt.Println("1. 구매")
// 		fmt.Println("2. 잔여 수량 확인")
// 		fmt.Println("3. 잔여 마일리지 확인")
// 		fmt.Println("4. 배송 상태 확인")
// 		fmt.Println("5. 장바구니 확인")
// 		fmt.Println("6. 프로그램 종료")
// 		fmt.Print("실행할 기능을 입력하시오 :")	
	
// 		fmt.Scanln(&menu)
// 		fmt.Println()
	
// 		if menu == 1 { // 물건 구매
// 			for {
// 				for i,v := range items {
// 					count := 1
// 					count += i
// 					fmt.Printf("물품%v : %v,가격: %v원, 잔여 수량 : %v\n",count, v.name, v.price , v.amount)
// 				}
// 				fmt.Print("구매할 물품을 선택하세요.")

// 				var choice int

// 				fmt.Scan(&choice)

// 				if choice > 6 {
// 					buyItem(choice,items,buyer)
// 					break;
// 				} else {
// 					fmt.Print("잘못된 입력입니다. 다시 입력해주세요.\n")
// 				}
				

// 				fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
// 				fmt.Scanln()
// 			}
// 		} else if menu == 2 { // 남은 수량 확인
// 			for _,v := range items {
// 				fmt.Printf("%v, 잔여수량 : %v\n",v.name,v.amount)
// 			}
// 			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
// 			fmt.Scanln()
// 		} else if menu == 3 { // 잔여 마일리지 확인
// 			fmt.Printf("현재 잔여 포인트는 %v점입니다.", buyer.point)
// 			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
// 			fmt.Scanln()
// 		} else if menu == 4 { // 배송 상태 확인
	
// 			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")	
// 			fmt.Scanln()
// 		} else if menu == 5 { // 장바구니 확인
	
// 			fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")	
// 			fmt.Scanln()
// 		} else if menu == 6 { // 프로그램 종료
// 			fmt.Println("프로그램을 종료합니다.")
	
// 			return // main함수 종료
// 		} else {
// 			fmt.Println("잘못된 입력입니다. 다시 입력해주세요.\n")
// 		}
// 	}	
// }