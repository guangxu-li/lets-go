package main

import (
	"fmt"
	"log"
)

func main() {
	var num float32
	fmt.Println("Please input one float number and hit the enter.")
	_, err := fmt.Scan(&num)

	if err == nil {
		fmt.Println("The truncated number is:", int(num))
	} else {
		fmt.Println("Please input a valid float number.")
		log.Fatal(err)
	}
}
