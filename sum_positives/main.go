package main

import (
	"fmt"
)

func main() {

	var numeros = [5]int{1, 2, -34, -56, 5}

	var res int
	res = PositiveSum(numeros)
	fmt.Println("Resultado: ", res)
}

func PositiveSum(numbers [5]int) int {

	sum := 0
	for _, value := range numbers {
		if value > 0 {
			fmt.Println("sum vale ", sum)
			fmt.Println("sumamos ", value)
			sum = sum + value
		}
	}
	return sum
}
