package model
//package main

import (
    "fmt"
    "../db"
    "github.com/go-pg/pg"
    "github.com/go-pg/pg/orm"
    "errors"
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
var roomsById map[int]*Room

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

func (room *Room) Exits(exits ...*Exit) Exits {
    return GroupExitsBy("From")[room.Id()]
}

func CreateRoom(base *Room_) *Room {
    base.Insert()
    room := &Room{Base: base, State:make(map[string]interface{})}
    roomsById[room.Id()] = room
    allRooms = append(allRooms, room)
    return room
}

func RoomById(id int) *Room {
    return roomsById[id]
}

func HomeRoom() *Room {
    return allRooms[0]
}

type myroomfunc func(*Room) bool
func (rooms *Rooms) Find(f myroomfunc) (*Room, error) {
    for _,e := range *rooms {
        if f(e) == true {
            return e, nil
        }
    }
    return nil, errors.New("value_not_found")
}

func init() {
    thefunc := func(room *Room) bool {
        return room.Id() == 2
    }

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
    roomsById = make(map[int]*Room)

    for i,_ := range allRooms_ {
        allRooms[i] = &Room{Base: &allRooms_[i], State:make(map[string]interface{})}
        roomsById[allRooms_[i].Id] = allRooms[i]
    }
    allRooms.Find(thefunc)
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

