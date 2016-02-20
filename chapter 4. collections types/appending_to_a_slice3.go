package main
import "fmt"

func main()  {
        cities := []string{"San Diego", "Mountain View"}
        otherCities := []string{"Santa Monica", "Venice"}
        cities = append(cities, otherCities...)
        fmt.Printf("%q", cities)
        // ["San Diego" "Mountain View" "Santa Monica" "Venice"]

}
