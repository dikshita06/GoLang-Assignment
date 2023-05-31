package main

import "fmt"

func main() {
	z := complex(4, 7)

	realPart := real(z)
	imaginaryPart := imag(z)

	sum := complex(4, 8) + complex(8, 3)
	difference := complex(5, 6) - complex(2, 3)
	product := complex(2, 3) * complex(4, 5)
	quotient := complex(10, 8) / complex(2, 2)

	fmt.Println("Complex number:", z)
	fmt.Println("Real part:", realPart)
	fmt.Println("Imaginary part:", imaginaryPart)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)

}
