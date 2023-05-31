package main
import "fmt"

func main() {
	var myString string
	myString = "Good morning!"

	name := "Dikshita"
	greeting := "Good morning, " + name + "!"

	firstChar := myString[0]

	length := len(myString)

	for _, char := range myString{
		fmt.Println("%c", char)
	}
	fmt.Println()

	fmt.Println(myString)
	fmt.Println(greeting)
	fmt.Println(firstChar)
	fmt.Println(length)
	
}