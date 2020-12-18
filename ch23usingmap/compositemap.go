package main

import "fmt"

func initMemberList() map[string]map[string]float64 {
	members := map[string]map[string]float64{
		"Tom": map[string]float64{
			"height": 175.4,
			"weight": 65,
		},
		"Jane": map[string]float64{
			"height": 168.3,
			"weight": 57.5,
		},
	}
	return members
}

func checkExist(members map[string]map[string]float64, member string, msg string) bool {
	fmt.Println("*** checkExist : ", msg)
	if _, ok := members[member]; ok {
		return true
	}
	return false
}

func printAllMembers(members map[string]map[string]float64) {
	fmt.Println("-----------------------")
	for key, value := range members {
		fmt.Println(key)
		fmt.Println(value["height"], value["weight"])
	}
	fmt.Println("-----------------------\n")
}

func addMember(list map[string]map[string]float64, name string, elem map[string]float64) {
	list[name] = elem
}

func deleteMember(list map[string]map[string]float64, name string) {
	delete(list, name)
}

func main() {
	members := initMemberList()
	printAllMembers(members)

	var newMember map[string]float64 = map[string]float64{
		"height": 180,
		"weight": 70,
	}

	addedName := "Willy"
	if ok := checkExist(members, addedName, "Before addMember "+addedName); ok {
		fmt.Println("exist")
	} else {
		fmt.Println("not exist")
	}
	addMember(members, addedName, newMember)

	if ok := checkExist(members, addedName, "after addMember "+addedName); ok {
		fmt.Println("exist")
	} else {
		fmt.Println("not exist")
	}

	printAllMembers(members)
	deletedName := "Jane"
	deleteMember(members, deletedName)
	if ok := checkExist(members, deletedName, "after deleteMember "+deletedName); ok {
		fmt.Println("exist")
	} else {
		fmt.Println("not exist")
	}
}
