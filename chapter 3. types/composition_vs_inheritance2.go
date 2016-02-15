package main

import "fmt"

type User struct {
        Id int
        Name, Location string
}

type Player struct {
        User
        GameId int
}

func main()  {
        p := Player{
                User{Id: 42, Name: "Matt", Location: "LA"},
                90404,
        }
        fmt.Printf(
                "Id: %d, Name: %s, Location: %s, Game id: %d\n",
                p.Id, p.Name, p.Location, p.GameId )
        // Directly set a field define on the Player struct
        p.Id = 11
        fmt.Printf("%+v", p)
}
