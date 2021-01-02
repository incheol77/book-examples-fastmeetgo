package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := make(map[string]interface{})
	data["name"] = "maria"
	data["age"] = 10

	doc, _ := json.MarshalIndent(data, "", " ")
	fmt.Println(string(doc))
}
