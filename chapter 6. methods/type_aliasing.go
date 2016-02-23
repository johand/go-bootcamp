package main

import (
	"fmt"
	"strings"
)

type MyStr string

func (s MyStr) UpperCase() string {
	return strings.ToUpper(string(s))
}

func main() {
	fmt.Println(MyStr("test").UpperCase())
}
