package main

import "fmt"

const MAX_NUMBER = 1000

func main() {
	multiples := []int{}

	for i := 1; i < MAX_NUMBER; i++ {
		if isMultipleOfThreeOrFive(i) {
			multiples = append(multiples, i)
		}
	}

	sum := 0
	for _, value := range multiples {
		sum = sum + value
	}

	fmt.Print(sum)
}

func isMultipleOfThreeOrFive(num int) bool {
	return num % 3 == 0 || num % 5 == 0
}
