package model

type Player struct {
    Id int
    Name string
    Location *Room
}
type Players []*Player
