package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var b float64 = 5.5
	var c complex128 = complex(3, 4)
	var name string = "Monica Gellar"
	var isTrue bool = true

	if a < 5 {
		fmt.Println("a is less than 5")
	} else {
		fmt.Println("a is greater than or equal to 5")
	}

	sum := a + int(b)
	product := b * float64(a)
	remainder := a % int(b)

	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)
	fmt.Println("Remainder:", remainder)

	switch a {
	case 1, 2, 3:
		fmt.Println("a is a small number")
	case 4, 5, 6:
		fmt.Println("a is a medium number")
	default:
		fmt.Println("a is a large number")

	}

	realPart := real(c)
	imaginaryPart := imag(c)

	fmt.Println("Real Part: %.2f\n", realPart)
	fmt.Println("Imaginary Part: %.2f\n", imaginaryPart)

	greeting := "Hello, " + name + "!"
	fmt.Println(greeting)
	fmt.Println("Length of name:", len(name))
	fmt.Println("first word of name:", string(name[0]))

	if isTrue {
		fmt.Println("The boolean value is true")
	} else {
		fmt.Println("The boolean value is false")
	}
}
