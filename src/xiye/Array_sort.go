package main

import "fmt"

func main() {
	var slice []int = []int{162, 127, 105, 87, 68, 54, 28, 18, 6, 3}
	fmt.Println(slice)
	fmt.Println(len(slice))
	insert := 30
	//var tmp int
	for i := 0; i < 10; i++ {
		if insert > slice[i] {
			slice = append(slice, 0)
			//fmt.Println(slice)
			for j := 9; j >= i; j-- {
				slice[j+1] = slice[j]
				//fmt.Println(slice)
			}
			slice[i] = insert
			break
		}
	}
	
	
	fmt.Println(slice)
}
