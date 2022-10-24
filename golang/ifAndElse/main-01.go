package test

import (
	"sort"
)

func Test1() {
	var arr2 = []int{}

	for i:=1; i<=100; i ++ {
		seven:= i*7
		naine:=i*9
		if seven < 100{
			arr2 = append(arr2, seven)
		}
		if naine < 100{
			if(naine !=63){
			arr2 = append(arr2, naine)
			}
		}
	}
	sort.Ints(arr2)

	for _,data := range arr2{
		print(data," ")
	}
}