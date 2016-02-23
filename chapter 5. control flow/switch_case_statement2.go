package main

import "fmt"

func main() {
	num := 3
	v := num % 2
	switch v {
	case 0:
		fmt.Println("even")
	case 3 - 2:
		fmt.Println("odd")
	}
}
