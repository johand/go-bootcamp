package main
import "fmt"

func main()  {
        a := [2]string{"Hello", "World!"}
        fmt.Println(a)
        // [Hello World!]
        fmt.Printf("%s\n", a)
        // [Hello World!]
        fmt.Printf("%q\n", a)
        // ["Hello" "World!"]
}
