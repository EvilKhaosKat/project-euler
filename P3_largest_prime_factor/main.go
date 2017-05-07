package main

import (
	"math"
	"fmt"
)

//using simple factors enumeration algorithm
//	check all natural numbers from 2 to sqrt of n
//      divide number over and over again
func main() {
	number := 600851475143

	primeFactors := []int{}
	var maxFactorToCheck int = getMaxFactorToCheck(number)
	fmt.Printf("maxFactorToCheck:%d\n", maxFactorToCheck)

	curNumber := number
	for {
		for divider := 2; divider < maxFactorToCheck; divider++ {
			if curNumber%divider == 0 {
				curNumber /= divider
				primeFactors = append(primeFactors, divider)

				break
			}
		}

		if curNumber == 1 {
			break
		}
	}

	fmt.Printf("primeFactors:%v\n", primeFactors)
	fmt.Printf("maxFactor:%d\n", getMaxValue(primeFactors))
}

func getMaxValue(values []int) int {
	maxValue := 0

	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

func getMaxFactorToCheck(num int) int {
	floatNum := float64(num)
	return int(math.Sqrt(floatNum))
}
