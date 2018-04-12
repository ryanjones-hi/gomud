package cmd

import "../model"
import "strconv"
import "fmt"

func Move(player *model.Player, params ...[]byte) string {
    exitId, err := strconv.Atoi(string(params[1]))
    if err != nil {
        panic(err)
    }
    exits := model.GroupExitsBy("From")[player.RoomId()]
   
   
    find := func(exit *model.Exit) bool { return exit.Id() == exitId}

    exit, err := exits.Find(find)

    if err != nil {
        fmt.Println(err)
        return "Exit does not exist here!" 
    }

    //currRoom := model.RoomById(player.RoomId())
    //nextRoom := model.RoomById(exit.To())
    player.SendMsg("foo")

    players := model.GroupPlayersBy("Room")[player.RoomId()]
    fmt.Println(exit.To())
    player.RoomId(exit.To())

    players.Broadcast(fmt.Sprintf("Player %v has left the building!",player.Id()))
    players = model.GroupPlayersBy("Room")[player.RoomId()]
    players.Broadcast(fmt.Sprintf("Player %v has entered the building!",player.Id()))

fmt.Println(players)
    
    return Look(player)

 
}
