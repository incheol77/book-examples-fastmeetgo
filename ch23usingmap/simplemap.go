package main

import "fmt"

func checkExist(map1 map[string]int, elem string) {
	if value, ok := map1[elem]; ok {
		fmt.Println(value, " exist.")
	} else {
		fmt.Println("not exist")
	}
}

func main() {
	var map1 map[string]int = make(map[string]int)
	a := [...]string{
		"aaa",
		"bbb",
		"ccc",
	}

	map1[a[0]] = 1
	map1[a[1]] = 2
	map1[a[2]] = 3

	for key, value := range map1 {
		fmt.Println(key, value)
	}

	checkExist(map1, a[0])
	delete(map1, "aaa")
	fmt.Println("After delete ", "aaa")
	checkExist(map1, a[0])
}
