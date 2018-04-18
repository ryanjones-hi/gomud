package cmd

import "../model"


func Logout(player *model.Player, params ...[]byte) {
    player.Logout()
}
