package main

import "fmt"

func StrRev(s string) string {
	rev := ""
	for i := StrLen(s) - 1; i >= 0; i-- {
		rev += string(s[i])
	}
	return rev
}

func main() {
	fmt.Println(StrRev("Hello World!"))
}
