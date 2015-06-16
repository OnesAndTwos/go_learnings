package main

import "fmt"

type tester func(int) bool

func isOdd(x int) bool {
	return x%2 == 1
}

func isEven(x int) bool {
	return x%2 == 0
}

func filter(xs []int, predicate tester) []int {
	result := []int{}

	for _, val := range xs {
		if predicate(val) {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	oddNumbers := filter(list, isOdd)
	evenNumbers := filter(list, isEven)

	fmt.Println(oddNumbers)
	fmt.Println(evenNumbers)
}
