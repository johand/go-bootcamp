package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	// same as "Bell Labs": Vertex{40.68433, -74.39967}
	"Google": {37.42202, -122.08408},
}

func main() {
	m["Splice"] = Vertex{34.05641, -118.48175}
	fmt.Println(m["Splice"])
	delete(m, "Splice")
	fmt.Printf("%v\n", m)
	name, ok := m["Splice"]
	fmt.Printf("key 'Splice' is present?: %t - value: %v\n", ok, name)
	name, ok = m["Google"]
	fmt.Printf("key 'Google' is present?: %t - value: %v\n", ok, name)
}
