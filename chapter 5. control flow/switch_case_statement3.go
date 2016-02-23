package main

import "fmt"

func main() {
	score := 7
	switch score {
	case 0, 1, 3:
		fmt.Println("Terrible")
	case 4, 5:
		fmt.Println("Mediocre")
	case 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Almost perfect")
	case 10:
		fmt.Println("hmm did ou cheat?")
	default:
		fmt.Println(score, " off the chart")
	}
}
