package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr2 = []int{}


	for i:=1; i<=100; i ++ {
		seven:= i*7
		naine:=i*9
		if seven < 100{
			arr2 = append(arr2, seven)
		}
		if naine < 100{
			arr2 = append(arr2, naine)
		}
	}
	sort.Ints(arr2)
	keys := make(map[int]bool) 

	ue := []int{} 
	fmt.Print((keys[7]))
	for _,value := range arr2{
		 if _,saveValue := keys[value]; !saveValue{
			  keys[value] = true 
			  ue = append(ue, value) 
			}
	} 
	for _,data := range ue{
		print(data," ")
	}


}