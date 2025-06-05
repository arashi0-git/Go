package main

import "fmt"

func Swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func main() {
	a := 1
	b := 2

	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}