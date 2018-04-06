package model
//package main

import (
    "fmt"

    "github.com/go-pg/pg"
    "github.com/go-pg/pg/orm"
)

type Room struct {
    Id int
    Name string
    Text string
}
type Rooms []*Room

func main() {
    Db := pg.Connect(&pg.Options{
        User:     "gomud",
        Password: "gomud",
        Database: "gomud",
    })
   err := Db.CreateTable((*Room)(nil),&orm.CreateTableOptions{})
   if err != nil {
       panic(err)
   }
 
   fmt.Println("foo")
}

