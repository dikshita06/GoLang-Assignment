package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{
		Name: "Dikshita Mate",
		Age:  25,
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error mareshelling JSON:", err)
		return
	}

	fmt.Println("Marshalled JSON data:", string(jsonData))

	var newPerson Person
	err = json.Unmarshal(jsonData, &newPerson)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println("New Person:", newPerson)

}
