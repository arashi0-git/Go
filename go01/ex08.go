package main

import "fmt"

func BasicAtoi(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		num = num * 10 + int(s[i] - '0')
	}
	return num
}

func main() {
	fmt.Println(BasicAtoi("000012345"))
	fmt.Println(BasicAtoi("0000"))

}
