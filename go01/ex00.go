package main

import "fmt"

func PointOne(nb *int) {
	*nb = 1
}

func main() {
	nb := 10
	PointOne(&nb)
	fmt.Println(nb)
}