package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(StringFormat32("01234567890123456789012345678901"))
	fmt.Println(StringFormat32("0123456789012345678901234567890123456"))
	fmt.Println(StringFormat32("01234567890123456789"))
	
	//fmt.Println(fmt.Sprintf("%032s", "1231"))
	s := []byte("hello")
	t1 := make([]byte, 32-len(s))
	for i := 0; i < len(t1); i++ {
		t1[i] = 97
	}
	t := append(s, t1...)
	fmt.Println(string(t))
}

func tt(){
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := data[8:]
	s2 := data[:5]
	copy(s2, s) // dst:s2, src:s
	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(data)
}

func StringFormat32(raw string) string {
	
	switch {
	case len(raw) == 32:
		return raw
	case len(raw) > 32:
		return fmt.Sprintf("%.32s",raw[:32])
	case len(raw) < 32:
		return fmt.Sprintf("%s%s",raw,strings.Repeat("_",32-len(raw)))
	default:
		/*should never meet here*/
		return raw
	}
}

