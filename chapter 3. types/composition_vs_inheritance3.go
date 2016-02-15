package main

import "fmt"

type User struct {
        Id int
        Name, Location string
}

func (u *User) Greetings() string {
        return fmt.Sprintf("Hi %s from %s",
                u.Name, u.Location)
}

type Player struct {
        User
        GameId int
}

func main()  {
        p := Player{}
        p.Id = 42
        p.Name = "Matt"
        p.Location = "LA"
        fmt.Println(p.Greetings())
}
