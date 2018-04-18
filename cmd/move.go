package cmd

import "../model"
import "strconv"
import "fmt"

func SysMove(player *model.Player, destination int) {
    players := model.GroupPlayersBy("Room")[player.RoomId()]
    player.RoomId(destination)

    players.Broadcast(fmt.Sprintf("Player %v has left the building!",player.Id()))
    players = model.GroupPlayersBy("Room")[player.RoomId()]
    Look(player)
    players.Broadcast(fmt.Sprintf("Player %v has entered the building!",player.Id()))

}

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
        player.SendMsg("That exit does not exist here!") 
        return
    }

    SysMove(player, exit.To())

}
