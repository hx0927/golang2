package main

import (
	"fmt"
)

func main() {
	str1 := "abc123"
	for index := range str1 {
		fmt.Printf("str1 -- index:%d, value:%d\n", index, str1[index])
	}

	str2 := "测试中文"
	for index := range str2 {
		fmt.Printf("str2 -- index:%d, value:%d\n", index, str2[index])
	}
	fmt.Printf("len(str2) = %d\n", len(str2))

	runesFromStr2 := []rune(str2)
	bytesFromStr2 := []byte(str2)
	fmt.Printf("len(runesFromStr2) = %d\n", len(runesFromStr2))

	
	fmt.Printf("len(bytesFromStr2) = %d\n", len(bytesFromStr2))
}
