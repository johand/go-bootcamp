package main
import "fmt"

func main()  {
        cities := make([]string, 3)
        cities[0] = "Santa Monica"
        cities[1] = "Venice"
        cities[2] = "Los Angeles"
        fmt.Printf("%q\n", cities)
        // ["Santa Monica" "Venice" "Los Angeles"]
}
