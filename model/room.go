package model
//package main

import (
//    "fmt"
    "../db"
    "errors"
)

//This base class mirrors the schema in the database
type Room_ struct {
    Id int
    Name string
    Desc string
}

type Room struct {
    Base  *Room_
    State map[string]interface{}
}
type Rooms []*Room
var AllRooms Rooms
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

func (room *Room) Desc(params ...string) string {
    if len(params) > 0 {
        room.Base.Desc = params[0]
    }
    return room.Base.Desc
}

func (room *Room) Exits(exits ...*Exit) Exits {
    return GroupExitsBy("From")[room.Id()]
}

func CreateRoom(base *Room_) *Room {
    base.Insert()
    room := &Room{Base: base, State:make(map[string]interface{})}
    roomsById[room.Id()] = room
    AllRooms = append(AllRooms, room)
    return room
}

func RoomById(id int) *Room {
    return roomsById[id]
}

func HomeRoom() *Room {
    return AllRooms[0]
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

    db.CreateTable((*Room_)(nil))
    var AllRooms_ []Room_
    err := db.Db.Model(&AllRooms_).Select()
    if err != nil {
        panic(err)
    }  
    if(len(AllRooms_) == 0) {
        room := Room_{Name:"Starting Room",Desc:"This is the room that all players start in!"}
        room.Insert()
        AllRooms_ = []Room_{room}
    }
    AllRooms = make(Rooms, len(AllRooms_))
    roomsById = make(map[int]*Room)

    for i,_ := range AllRooms_ {
        AllRooms[i] = &Room{Base: &AllRooms_[i], State:make(map[string]interface{})}
        roomsById[AllRooms_[i].Id] = AllRooms[i]
    }
}
