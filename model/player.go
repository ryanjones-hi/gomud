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
    Room *Room
}

type Player struct {
    Base *Player_
    State map[string]interface{}
}
type Players []*Player

func (player *Player_) Insert() {
    if err := db.Db.Insert(player); err != nil {
        panic(err)
    }    
}

func (player *Player) Id() int{
    return player.Base.Id
}

func (player *Player) Room(params ...*Room) *Room {
    if len(params) > 0 {
        player.Base.Room = params[0]
    }
    return player.Base.Room
}

func CreatePlayer(base *Player_) *Player {
    if base.Room == nil {
        base.Room = HomeRoom()
    }
    base.Insert()
    return &Player{Base: base, State: make(map[string]interface{})}
}

func BuildPlayerTable() {
    Db := pg.Connect(&pg.Options{
        User:     "gomud",
        Password: "gomud",
        Database: "gomud",
    })
   err := Db.CreateTable((*Player_)(nil),&orm.CreateTableOptions{})
   if err != nil {
       panic(err)
   }

   fmt.Println("foo")
}
