package pkg

import (
	"encoding/json"
	"fmt"
)

func Example() {
	config := &FetchConfig{
		method: "GET",
		body:   nil,
	}
	resp, err := Fetch("https://jsonplaceholder.typicode.com/todos/1", *config)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(resp.Body)
	var t struct {
		UserId    int    `json:"userId"`
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	err = decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Title)
}
