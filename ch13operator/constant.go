package main

import "fmt"

func main() {
	// const (
	// 	SUNDAY    = 0
	// 	MONDAY    = 1
	// 	TUESDAY   = 2
	// 	WEDNESDAY = 3
	// 	THURSDAY  = 4
	// 	FRIDAY    = 5
	// 	SATURDAY  = 6
	// )

	// const (
	// 	SUNDAY = iota
	// 	MONDAY
	// 	TUESDAY
	// 	WEDNESDAY
	// 	THURSDAY
	// 	FRIDAY
	// 	SATURDAY
	// )

	const (
		SUNDAY = 1 << iota
		MONDAY
		TUESDAY
		WEDNESDAY
		THURSDAY
		FRIDAY
		SATURDAY
	)

	fmt.Println(SUNDAY)
	fmt.Println(MONDAY)
	fmt.Println(TUESDAY)
	fmt.Println(WEDNESDAY)
	fmt.Println(THURSDAY)
	fmt.Println(FRIDAY)
	fmt.Println(SATURDAY)
}
