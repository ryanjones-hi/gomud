// Copyright 2013 The Goril:la WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//Vil: I'm thinking that this file will mostly concern outward communication.
package conn

import (
    "fmt"
    //"ourchat/cmd"
    //"github.com/go-pg/pg"
    //"github.com/go-pg/pg/orm"
    //"bytes"
    //"ourchat/model"
    //"ourchat/db"
)
// hub maintains the set of active clients and broadcasts messages to the
// clients.

type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

//This map associates types(e.g. room) with players(clients) associated with the type.
//For example, one string might be models.Room:4, the type + the id.
type ClientList []*Client
var groups map[string]ClientList

func NewHub() *Hub {
        groups = make(map[string]ClientList)
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}


func GetGroup(t interface{}, id int)(group ClientList) {
        key := fmt.Sprintf("%T:%d",t,id)

        if val, ok := groups[key]; ok {
            return val
        }
        return ClientList{}
}

func AddClient(t interface{}, id int, c *Client) {
        key := fmt.Sprintf("%T:%d",t,id)

        if val, ok := groups[key]; ok {
            groups[key] = append(val, c)
        } else {
            groups[key] = ClientList{c}
        }
        fmt.Println(groups[key])
}

func (h *Hub) Run() {
//    db.InitDb();

//    room := model.Room{
//        Name: "Foyer",
//        Text: "A cozy foyer.",
//    }
//    mydb := db.Db
//    fmt.Println(mydb)
//    err := mydb.Insert(&room)
//    if err != nil {
//        panic(err)
//    }
//    fmt.Println(room)
//
//    var n int
//    _, err = db.Db.QueryOne(pg.Scan(&n), "SELECT 1")
//    if err != nil {
//        panic(err)
//    }
//    fmt.Println(n)
//
//    err = db.Db.CreateTable(&model.Room{}, &orm.CreateTableOptions{
//        Temp: false, // create temp table
//    })
//
//    if err != nil {
//        //panic(err)
//        fmt.Println(err)
//    }
//
//    var info []struct {
//        ColumnName string
//        DataType string
//    }
//
//    _, err = db.Db.Query(&info, `
//        SELECT column_name, data_type
//        FROM information_schema.columns
//        WHERE table_name = 'hello_models'
//        `)
//    if err != nil {
//        panic(err)
//    }
//    fmt.Println(info)

	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}

}
