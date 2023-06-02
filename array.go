package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	person := Person{
		Name: "Monica Gellar",
		Age:  28,
	}

	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)

	ages := map[string]int{
		"Monica": 28,
		"Rachel": 29,
	}

	for key, value := range ages {
		fmt.Println(key, ":", value)
	}

	sum := add(5, 4)
	fmt.Println("Sum:", sum)
}
func add(a, b int) int {
	return a + b
}
