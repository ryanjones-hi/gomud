package db

import (
    "github.com/go-pg/pg"
    "github.com/go-pg/pg/orm"
    "fmt"
)

var Db *pg.DB

func CreateTable(schema interface{}) {
    Db.CreateTable(schema,&orm.CreateTableOptions{})
}

func InitDb() {

    Db = pg.Connect(&pg.Options{
        User:     "gomud",
        Password: "gomud",
        Database: "gomud",
    })
    fmt.Println(Db)
}

func init() {
    InitDb() 
}
