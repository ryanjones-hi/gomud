package db

import (
    "github.com/go-pg/pg"
    "fmt"
)

var Db *pg.DB

func InitDb() {

    Db = pg.Connect(&pg.Options{
        User:     "django",
        Password: "django",
        Database: "django_project",
    })
    fmt.Println(Db)
}
