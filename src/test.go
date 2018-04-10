package main

import (
	"fmt"
)

func main() {
	//fmt.Printf("haha", 3.14)
	//fmt.Println("haha", 3.14)

	map1()
	//slice1()
	//slice2()
}

func map1() {
	var colors map[string]string = make(map[string]string)
	colors["red"] = "666"
	fmt.Println(`colors["blue"]`, colors["blue"])

	value, exists := colors["blue"]
	if exists {
		fmt.Println(value)
	}

	//value,ok:=colors["red"]
	//if ok{
	//	fmt.Println(value)
	//}

	v := colors["red"]
	if v != "" {
		fmt.Println(v)
	}

}

func slice2() {
	slice := []int{10, 20, 30, 40}
	for i, v := range slice {
		fmt.Printf("Value: %d	value-Addr: %X Element: %X\n", v, &v, &slice[i])
	}
}

func slice1() {
	slice := []int{1, 2, 3, 4, 5}
	nslice := slice[1:3]
	fmt.Println(nslice)
	fmt.Println(len(nslice), cap(nslice))

	return

	nslice[1] = 999
	fmt.Println(slice)
	fmt.Println(nslice)

	nslice = append(nslice, 6)
	nslice = append(nslice, 7)
	nslice = append(nslice, 8)
	nslice = append(nslice, 9)
	nslice = append(nslice, 10)

	fmt.Println(slice)
	fmt.Println(nslice)

}
