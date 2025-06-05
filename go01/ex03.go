package main

import "fmt"

func UltimateDivMod(a *int, b *int) {
	tmpa := *a
	tmpb := *b
	*a = tmpa / tmpb
	*b = tmpa % tmpb
}

func main() {
	a := 13
	b := 2

	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
