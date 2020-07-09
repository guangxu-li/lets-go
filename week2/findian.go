package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Please input a string.")

	var str string
	fmt.Scan(&str)
	str = strings.ToLower(str)

	chars := []rune(str)
	if chars[0] == 'i' && chars[len(str)-1] == 'n' && strings.ContainsRune(str, 'a') {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
