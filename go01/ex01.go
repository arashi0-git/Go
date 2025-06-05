package main

import "fmt"

func UltimatePointOne(n ***int) {
	***n = 1
}

func main() {
	a := 0
	fmt.Println(a)
	fmt.Println(&a)
	b := &a
	n := &b
	UltimatePointOne(&n)
	fmt.Println(a)
	fmt.Println(&a)
}