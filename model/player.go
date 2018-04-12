package model
//package main

import (
    "fmt"
    "../db"
    "github.com/go-pg/pg"
    "github.com/go-pg/pg/orm"
)

//This is the portion of our player that is backed up in the database
type Player_ struct {
    Id int
    RoomId int
    Name string
}

type Player struct {
    Base *Player_
    Send chan []byte
    State map[string]interface{}
}
type Players []*Player
var allPlayers Players
var playersGroupedBy map[string]map[int]Players

func (player *Player_) Insert() {
    if err := db.Db.Insert(player); err != nil {
        panic(err)
    }    
}

func (player *Player) Id() int{
    return player.Base.Id
}

func (player *Player) SendMsg(msg string) {
    fmt.Println(msg)
    //fmt.Println("Player:",(*player).Base)
    //fmt.Println("SendMsg player.Send:", (*player).Send)
    player.Send <- []byte(msg)
}

func (player *Player) RoomId(params ...int) int {
    if len(params) > 0 {
        player.Base.RoomId = params[0]
    }
    return player.Base.RoomId
}

func (player *Player) Name(params ...string) string {
    if len(params) > 0 {
        player.Base.Name = params[0]
    }
    return player.Base.Name
}

func CreatePlayer(base *Player_) *Player {
    if base.RoomId == 0 {
        base.RoomId = HomeRoom().Id()
    }
    base.Insert()
    newPlayer := &Player{Base: base, State: make(map[string]interface{})}
    allPlayers = append(allPlayers,newPlayer)
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
            //return allExits
    }
}

func init() {
    Db := pg.Connect(&pg.Options{
        User: "gomud",
        Password: "gomud",
        Database: "gomud",
    })
    //TODO: Log the error below(Which we're probably fine with)
    Db.CreateTable((*Player_)(nil),&orm.CreateTableOptions{})

    var allPlayers_ []Player_
    err := Db.Model(&allPlayers_).Select()
    if err!=nil {
        panic(err)
    }
    allPlayers = make(Players,0)

    groupedbyroom := make(map[int]Players)
    //for i,_ := range allPlayers_ {
    //    newPlayer := &Player{Base: &allPlayers_[i], State:make(map[string]interface{})}
    //    allPlayers[i] = newPlayer
    //    if _, ok := groupedbyroom[newPlayer.RoomId()]; !ok {
    //        groupedbyroom[newPlayer.RoomId()] = Players{newPlayer}
    //    } else {
    //        groupedbyroom[newPlayer.RoomId()] = append(groupedbyroom[newPlayer.RoomId()],newPlayer)
    //    }
    //}

    playersGroupedBy = make(map[string]map[int]Players)
    playersGroupedBy["Room"] = groupedbyroom
    fmt.Println(allPlayers_)
}
