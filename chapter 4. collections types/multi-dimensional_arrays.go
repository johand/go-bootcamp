package main

import "fmt"

func main()  {
        var a [2] [3]string
        for i := 0; i < 2; i++ {
                for j := 0; j < 3; j++ {
                        a[i] [j] = fmt.Sprintf("row %d - column %d", i+1, j+1)
                }
        }
        fmt.Printf("%q\n", a)
        // [["row 1 - column 1" "row 1 - column 2" "row 1 - column 3"]
        // ["row 2 - column 1" "row 2 - column 2" "row 2 - column 3"]]
}
