package model
//package main

import (
    "fmt"
    "../db"
    "errors"
)

//This is the portion of our player that is backed up in the database
type Player_ struct {
    Id int
    RoomId int
    Name string
    Pass []byte
    Desc string
}

type Player struct {
    Base *Player_
    Send chan []byte
    State map[string]interface{}
}
type Players []*Player
var AllPlayers Players
var playersGroupedBy map[string]map[int]Players

func (player *Player_) Insert() {
    if err := db.Db.Insert(player); err != nil {
        panic(err)
    }    
}

func (player *Player) Id() int{
    return player.Base.Id
}

func (player *Player) Login(id int) error {

    if _,_,err := AllPlayers.Find(func(_player *Player)bool { return _player.Id() == id }); err==nil {
        return errors.New("Player is already logged in!")
    }
    playersGroupedBy["Room"][player.Base.RoomId] = append(playersGroupedBy["Room"][player.Base.RoomId], player)
    AllPlayers = append(AllPlayers, player)
    fmt.Println(AllPlayers)
    return nil
}

func (player *Player) Logout() {
    var err error
    playersGroupedBy["Room"][player.Base.RoomId],err = playersGroupedBy["Room"][player.Base.RoomId].Remove(player)
    if err != nil {
        panic(err)
    }
    AllPlayers,err = AllPlayers.Remove(player)
    if err != nil {
        panic(err)
    }
    player.Base = nil
    player.SendMsg("You have been logged out!")
    fmt.Println(AllPlayers)
}

func (player *Player) SendMsg(msg string) {
    player.Send <- []byte(msg)
}

func (player *Player) RoomId(params ...int) int {
    if len(params) > 0 {
        players := playersGroupedBy["Room"][player.Base.RoomId]
        i,e,_ := players.Find(func(_player *Player)bool { return player.Id() == _player.Id() })
        playersGroupedBy["Room"][player.Base.RoomId] = append(players[:i],players[i+1:]...)
        playersGroupedBy["Room"][params[0]] = append(playersGroupedBy["Room"][params[0]],e)
        player.Base.RoomId = params[0]
        db.Db.Update(player.Base)
    }
    return player.Base.RoomId
}

func (player *Player) Name(params ...string) string {
    if len(params) > 0 {
        player.Base.Name = params[0]
    }
    return player.Base.Name
}

func (player *Player) Desc(params ...string) string {
    if len(params) > 0 {
        player.Base.Desc = params[0]
        db.Db.Update(player.Base)
    }
    return player.Base.Desc
}


func CreatePlayer(base *Player_) *Player {
    if base.RoomId == 0 {
        base.RoomId = HomeRoom().Id()
    }
    base.Insert()
    newPlayer := &Player{Base: base, State: make(map[string]interface{})}
    AllPlayers = append(AllPlayers,newPlayer)
    groupedbyroom := playersGroupedBy["Room"]
    if _, ok := groupedbyroom[newPlayer.RoomId()]; !ok {
        groupedbyroom[newPlayer.RoomId()] = Players{newPlayer}
    } else {
        groupedbyroom[newPlayer.RoomId()] = append(groupedbyroom[newPlayer.RoomId()],newPlayer)
    }
    return newPlayer
}

func (players *Players) Broadcast(msg string) {
    fmt.Println("From players.Broadcast:",players)
    for _,p := range *players {
        p.SendMsg(msg)
    }
}

func GroupPlayersBy(prop string) map[int]Players {
    switch prop {
        case "Room":
            return playersGroupedBy["Room"]
        default:
            panic("Property not found!")
    }
}

type myplayerfunc func(*Player) bool
func (players *Players) Find(f myplayerfunc) (int, *Player, error) {
    for i,e := range *players{
        if f(e) == true {
            return i, e, nil
        }
    }
    return 0, nil, errors.New("item_not_found")
}

func (players Players) Remove(player *Player) (Players, error) {
    for i,e := range players{
        if e == player {
            return append(players[:i],players[i+1:]...), nil
        }
    }
    return players, errors.New("item_not_found") 
}

func init() {
    db.CreateTable((*Player_)(nil))

    var AllPlayers_ []Player_
    err := db.Db.Model(&AllPlayers_).Select()
    if err!=nil {
        panic(err)
    }
    AllPlayers = make(Players,0)

    groupedbyroom := make(map[int]Players)

    playersGroupedBy = make(map[string]map[int]Players)
    playersGroupedBy["Room"] = groupedbyroom
}
