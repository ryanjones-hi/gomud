package cmd

import "../model"
import "fmt"
import "strings"
import "strconv"

func LookExits(player *model.Player) string {
    exits := model.GroupExitsBy("From")[player.RoomId()]
    str := []string{}
    for _,e := range exits {
		str = append(str, fmt.Sprintf("%v(%v)",e.Name(),e.Id()))
    }

    return "Exits: " + strings.Join(str,", ")
}

func LookPlayers(player *model.Player) string {
    players := model.GroupPlayersBy("Room")[player.RoomId()]
    str := []string{}
    for _,e := range players {
        str = append(str, fmt.Sprintf("%v",e.Id()))
    }

    return "Players: " + strings.Join(str,", ")
}

func Look(player *model.Player, params ...[]byte) {
    if len(params) <= 1 {
        room := model.RoomById(player.RoomId())
        player.SendMsg(fmt.Sprintf("%v(%v)\n%v\n%v\n%v",room.Name(),room.Id(),room.Desc(),LookExits(player),LookPlayers(player)))
        return
    }

    fmt.Println("lookParams",params)
    switch string(params[1]) {
        case "exits":
            exits := model.GroupExitsBy("From")[player.RoomId()]
            str := []string{}
            for _,e := range exits {
			str = append(str, fmt.Sprintf("%v(%v)",e.Desc(),e.Id()))
            }
            player.SendMsg("Exits: " + strings.Join(str,", "))
           
        default:
            players := model.GroupPlayersBy("Room")[player.RoomId()]
            playerId, err := strconv.Atoi(string(params[1]))
            if err != nil {
                player.SendMsg("Invalid paramters!")
                return
            }

            _,target,err := players.Find(func(_player *model.Player)bool {return _player.Id() == playerId}) 

            if err != nil {
                player.SendMsg("That player is not here!")
                return
            }

            player.SendMsg(target.Desc())
            return        
    }
    player.SendMsg("No return")
}
