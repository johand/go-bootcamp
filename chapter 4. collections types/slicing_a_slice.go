package main
import "fmt"

func main()  {
        mySlice := []int{2, 3, 5, 7, 11, 13}
        fmt.Println(mySlice)
        // [2 3 5 7 11 13]
        fmt.Println(mySlice[1:4])
        // [3, 5, 7]

        // missing low index implies 0
        fmt.Println(mySlice[:3])
        // [2, 3, 5]

        // missing high index implies len (s)
        fmt.Println(mySlice[4:])
        // [11 13]
}
