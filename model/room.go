package model
//package main

import (
    "fmt"
    "../db"
    "github.com/go-pg/pg"
    "github.com/go-pg/pg/orm"
)

//This base class mirrors the schema in the database
type Room_ struct {
    Id int
    Name string
    Text string
}

type Room struct {
    Base  *Room_
    State map[string]interface{}
}
type Rooms []*Room
var allRooms Rooms

func (room *Room_) Insert() {
      if err := db.Db.Insert(room); err != nil {
      panic(err)
    }
}

func (room *Room) Insert() {
    room.Base.Insert()
}

func (room *Room) Id() int {
    return room.Base.Id
}

func (room *Room) Name(params ...string) string {
    if len(params) > 0 {
        room.Base.Name = params[0]
    }
    return room.Base.Name
}

func (room *Room) Text(params ...string) string {
    if len(params) > 0 {
        room.Base.Text = params[0]
    }
    return room.Base.Text
}

func CreateRoom(base *Room_) *Room {
    base.Insert()
    return &Room{Base: base, State:make(map[string]interface{})}
}

func HomeRoom() *Room {
    return allRooms[0]
}

func init() {
    Db := pg.Connect(&pg.Options{
        User:     "gomud",
        Password: "gomud",
        Database: "gomud",
    })
    Db.CreateTable((*Room_)(nil),&orm.CreateTableOptions{})
    var allRooms_ []Room_
    err := Db.Model(&allRooms_).Select()
    if err != nil {
        panic(err)
    }  
    allRooms = make(Rooms, len(allRooms_))

    for i,base := range allRooms_ {
        allRooms[i] = &Room{Base: &base, State:make(map[string]interface{})}
    }
    fmt.Println(allRooms)
}

func BuildRoomTable() {
    Db := pg.Connect(&pg.Options{
        User:     "gomud",
        Password: "gomud",
        Database: "gomud",
    })
   err := Db.CreateTable((*Room_)(nil),&orm.CreateTableOptions{})
   if err != nil {
       panic(err)
   }
 
   fmt.Println("foo")
}

