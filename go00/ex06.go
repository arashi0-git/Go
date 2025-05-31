package main

import "fmt"

var first = true

func printCombination(prefix string, start, n, digits int) {
	if n == 0 {
		if !first {
			fmt.Print(", ")
		}
		fmt.Print(prefix)
		first = false
		return
	}

	for i := start; i <= 9; i++ {
		newPrefix := prefix + fmt.Sprintf("%d", i)
		printCombination(newPrefix, i+1, n-1, digits)
	}
}

func PrintCombN(n int) {
	if n <= 0 || n > 9 {
		return
	}
	first = true
	printCombination("", 0, n, n)
	fmt.Println()
}

func main() {
	PrintCombN(1)
	PrintCombN(3)
	PrintCombN(9)
}