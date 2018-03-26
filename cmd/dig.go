package cmd

import "github.com/go-pg/pg"
import "ourchat/model"

func Dig(db *pg.DB, params ...[]byte) {

//    cmd := params[0]
    name := params[1]
    text := params[2]

    room := model.Room{
        Name: string(name),
        Text: string(text),
    }

    if err := db.Insert(&room); err != nil {
        panic(err)
    }
}
