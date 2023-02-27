package main

import (
	"fmt"
)

func inputSubNum() (int, error) {		
	var num int
	
	fmt.Scanln(num)
	
	if num <= 0{
		err := fmt.Errorf("잘못된 과목 수입니다.")

		return 0, err
	}
	
	return num , nil
}

func average(num int) (float64 , error) {
	var score, total int
	
	for i := 0; i < num; i++ {
		fmt.Scanln(&score)
		if score < 0 || score > 100{
			err := fmt.Errorf("잘못된 점수입니다.")
	
			return 0, err
		}
		total += score
	}	
	
	avg := float64(total)/ float64(num)
	
	return avg , nil
}

func main() {	
	defer func ()  {
		r:= recover()
		if r != nil{
			fmt.Print(r)

		}

	}()
	
	
	 num, error := inputSubNum()
	
	if error != nil {
		panic(error)
	}
	
	 result, error := average(num)

	 if error != nil {
		panic(error)
	}
	
	
	fmt.Printf("%.1f",result)	
}