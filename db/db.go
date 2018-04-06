package db

import (
    "github.com/go-pg/pg"
    "fmt"
)

var Db *pg.DB

func InitDb() {

    Db = pg.Connect(&pg.Options{
        User:     "gomud",
        Password: "gomud",
        Database: "gomud",
    })
    fmt.Println(Db)
}
