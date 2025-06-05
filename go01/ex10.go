package main

import "fmt"

func StrLen(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count++
	}
	return count
}

func Atoi(s string) int {
	num := 0
	flag := 1
	startIndex := 0

	// 空文字列チェック
	if StrLen(s) == 0 {
		return 0
	}

	// 符号の処理（最初の文字のみ）
	if s[0] == '-' {
		flag = -1
		startIndex = 1
	} else if s[0] == '+' {
		flag = 1
		startIndex = 1
	}

	// 符号の後に数字があるかチェック
	if startIndex >= StrLen(s) {
		return 0
	}

	// 数字の処理
	for i := startIndex; i < StrLen(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		num = num*10 + int(s[i]-'0')
	}

	return num * flag
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("-12345"))
	fmt.Println(Atoi("+12345"))
	fmt.Println(Atoi("++12345"))
	fmt.Println(Atoi("--12345"))
	fmt.Println(Atoi("12345-"))
	fmt.Println(Atoi("12345+"))
	fmt.Println(Atoi("12345+12345"))
	fmt.Println(Atoi("12345-12345"))
	fmt.Println(Atoi("12345+12345"))
}