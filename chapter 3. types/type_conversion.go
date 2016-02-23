package main

import "fmt"

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

func main()  {
        fmt.Println(i)
        fmt.Println(f)
        fmt.Println(u)
}
