package main
import "fmt"

func main() {
        cities := []string{
                "Santa Monica",
                "San Diego",
                "San Francisco",
        }
        fmt.Println(len(cities))
        // 3
        countries := make([]string, 42)
        countries[0] = "NY"
        countries[15] = "Nevada"
        fmt.Println(len(countries))
        // 42
}
