package main

import (
	"fmt"
)

func main() {
	str1 := "abc123"
	for index := range str1 {
		fmt.Printf("str1 -- index:%d, value:%d\n", index, str1[index])
	}


}
