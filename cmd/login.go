package cmd

import "../model"
import "golang.org/x/crypto/bcrypt"
import "../db"

func Login(player *model.Player, params ...[]byte) {
    if len(params) != 3 {
        player.SendMsg("Format: login name password")
        return
    }
    name := string(params[1])
    password := params[2]
    player_ := &model.Player_{Name:name}
    if err := db.Db.Model(player_).Where("name = ?", name).Select(); err != nil {
        player.SendMsg(err.Error())
        return
    } else {
        if err = bcrypt.CompareHashAndPassword(player_.Pass, password); err != nil {
            player.SendMsg("Bad username/password")
            return
	}
    }

    player.Base = player_
    if err := player.Login(player_.Id); err != nil {
        player.SendMsg("That player is already logged in!")
        player.Base = nil
        return
    }
    
    player.SendMsg("Logged in!")
    Look(player)
}
