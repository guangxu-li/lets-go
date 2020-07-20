package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 1, 3)
	first := true

	for {
		fmt.Println("Please input an Integer or \"X\" to quit")

		var input string
		fmt.Scan(&input)

		if input == "X" {
			break
		}

		if num, err := strconv.Atoi(input); err == nil {
			if first {
				sli[0] = num
				first = false
			} else {
				sli = append(sli, num)
			}

			sort.Ints(sli)
			fmt.Println(sli)
		} else {
			fmt.Println("The input is not valid.")
		}
	}
}
