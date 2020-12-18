package main

import "fmt"

func add(arr ...int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func sumArr(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func main() {

	sum := add(1, 2, 3, 4, 5)
	fmt.Println(sum)

	sum = add(1, 2, 3)
	fmt.Println(sum)

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(sumArr(arr))
}
