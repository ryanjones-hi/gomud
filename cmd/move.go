package cmd

import "../model"
import "strconv"
import "fmt"

func Move(player *model.Player, params ...[]byte) {
    exitId, err := strconv.Atoi(string(params[1]))
    if err != nil {
        player.SendMsg("Invalid parameter!")
        return
    }
    exits := model.GroupExitsBy("From")[player.RoomId()]
   
   
    find := func(exit *model.Exit) bool { return exit.Id() == exitId}

    exit, err := exits.Find(find)

    if err != nil {
        player.SendMsg("Exit does not exist here!") 
        return
    }

    player.SendMsg("foo")

    players := model.GroupPlayersBy("Room")[player.RoomId()]
    player.RoomId(exit.To())

    players.Broadcast(fmt.Sprintf("Player %v has left the building!",player.Id()))
    players = model.GroupPlayersBy("Room")[player.RoomId()]
    players.Broadcast(fmt.Sprintf("Player %v has entered the building!",player.Id()))

    Look(player)
}
