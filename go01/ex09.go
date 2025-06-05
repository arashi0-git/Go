package main

import "fmt"

func StrLen(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count++
	}
	return count
}

func BasicAtoi2(s string) int {
	num := 0
	for i := 0; i < StrLen(s); i++ {
		if (s[i] < '0' || s[i] > '9') {
			return 0
		}
		num = num * 10 + int(s[i] - '0')
	}
	return num
}

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}