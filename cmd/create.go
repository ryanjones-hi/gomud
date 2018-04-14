package cmd

import "../model"
import "golang.org/x/crypto/bcrypt"
import "../db"

func Create(player *model.Player, params ...[]byte) {
    if len(params) != 3 {
        player.SendMsg("Format: create name password")
        return
    }
    name := string(params[1])
    password := params[2]
    hashed, err := bcrypt.GenerateFromPassword(password, 8) 
    player_ := &model.Player_{Name:name}
    if err = db.Db.Model(player_).Where("name = ?", name).Select(); err == nil {
        player.SendMsg("Player already exists!")
        return
    }

    player_ = &model.Player_{Name:name,Pass:hashed,RoomId:model.HomeRoom().Id()}
    db.Db.Insert(player_)
    player.SendMsg("Player created! Please login with LOGIN command")
}
