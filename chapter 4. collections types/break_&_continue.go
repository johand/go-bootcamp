package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
		if pow[i] >= 16 {
			break
		}
	}
	fmt.Println(pow)
}



