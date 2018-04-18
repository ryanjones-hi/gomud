package cmd

import "../model"
import "fmt"

func Say(player *model.Player, params ...[]byte) {
    if len(params) < 2 {
        player.SendMsg("SAY requires a message")
    }

    players := model.GroupPlayersBy("Room")[player.RoomId()]

    
    msg := fmt.Sprintf("%v says, \"%v\"",player.Name(),string(params[1]))
//    msg := fmt.Fprintf("%v says, %v","whoa","hey")
    players.Broadcast(msg)
}

