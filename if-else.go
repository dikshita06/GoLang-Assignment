package main
import "fmt"

func main() {
	score := 80

	if score >= 90 {
		fmt.Println("You got an A grade.")
	} else if score >= 80 {
		fmt.Println("You got an B grade.")
	} else if score >= 70 {
		fmt.Println("You got an C grade.")
	} else {
		fmt.Println("You need to improve your score.")
	}
}