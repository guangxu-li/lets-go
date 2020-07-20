package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 0, 3)

	for {
		fmt.Println("Please input an Integer or \"X\" to quit")

		var input string
		fmt.Scan(&input)

		if input == "X" {
			break
		}

		if num, err := strconv.Atoi(input); err == nil {
			sli = append(sli, num)

			sort.Ints(sli)
			fmt.Println(sli)
		} else {
			fmt.Println("The input is not valid.")
		}
	}
}
