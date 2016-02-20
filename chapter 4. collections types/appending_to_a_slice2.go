package main
import "fmt"

func main()  {
        cities := []string{}
        cities = append(cities, "San Diego", "Mountain View")
        fmt.Printf("%q", cities)
}
