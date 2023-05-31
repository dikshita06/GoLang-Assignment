package main

import "fmt"

func main() {
	var x float32
	var y float64

	x = 3.14
	y = 2.1865

	result := x + float32(y)
	difference := x - float32(y)
	product := x * float32(y)

	fmt.Println("value: %.2f\n", result)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)

}
