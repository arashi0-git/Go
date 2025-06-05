package main

import "fmt"

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table) - 1; j++ {
			if (table[j] > table[j + 1]) {
				table[j], table[j + 1] = table[j + 1], table[j]
			}
		}
	}
}

func main() {
	s := []int{4,3,5,1,6,0}
	SortIntegerTable(s)
	fmt.Println(s)
}
