package main

import "fmt"

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]))
	}
	fmt.Println()
}

func main() {
	PrintStr("Hello World!")
}