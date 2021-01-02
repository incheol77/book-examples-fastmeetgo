package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	Name  string
	Email string
}

type Comment struct {
	Id      uint64
	Author  Author
	Content string
}

type Article struct {
	Id         uint64
	Title      string
	Author     Author
	Content    string
	Recommends []string
	Comments   []Comment
}

func main() {
	data := make([]Article, 1)

	data[0].Id = 1
	data[0].Title = "Hello world"
	data[0].Author.Name = "Marial"
	data[0].Author.Email = "marial@email.com"
	data[0].Content = "Hello~"
	data[0].Recommends = []string{"John", "Andrew"}
	data[0].Comments = make([]Comment, 1)
	data[0].Comments[0].Id = 1
	data[0].Comments[0].Author.Name = "Andrew"
	data[0].Comments[0].Author.Email = "andrew@email.com"
	data[0].Comments[0].Content = "Hello Maria"

	doc, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(doc))
}
