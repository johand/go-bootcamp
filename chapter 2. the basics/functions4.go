package main

import "fmt"

func location(city string) (region, country string)  {
        switch city {
        case "Los Angeles", "LA", "Santa Monica":
                region, country = "California", "North America"
        case "New York", "NYC":
                region, country = "New York", "North America"
        default:
                region, country = "Unknown", "Unknown"
        }
        return
}

func main()  {
        region, country := location("Santa Monica")
        fmt.Printf("Matt lives in %s, %s", region, country)
}
