package cmd

import "../model"
import "fmt"
import "strconv"

func Whisper(player *model.Player, params ...[]byte) {
    if len(params) < 3 {
        player.SendMsg("Format: WHISPER recipient message")
        return
    }


    recipientId, err := strconv.Atoi(string(params[1]))
    if err != nil {
        player.SendMsg("Invalid parameter for recipient id!")
        return
    }

    //players := model.GroupPlayersBy("Room")[player.RoomId()]
    msg := fmt.Sprintf("%v whispers: \"%v\"",player.Name(),string(params[2]))
    _,recipient,err := model.AllPlayers.Find(func(_player *model.Player)bool { return _player.Id() == recipientId})

    if err != nil {
        player.SendMsg("Player not found!")
        return
    }
    
    recipient.SendMsg(msg)
    player.SendMsg(msg)

    
}

