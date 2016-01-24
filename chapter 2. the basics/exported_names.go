package main

import (
        "fmt"
        "math"
)

// cannot refer to unexported name math.pi
// undefined: math.pi
// func main()  {
//         fmt.Println(math.pi)
// }

func main()  {
        fmt.Println(math.Pi)
}
