package cmd

import "../model"
import "fmt"
import "strings"

func LookExits(player *model.Player) string {
    exits := model.GroupExitsBy("From")[player.RoomId()]
    str := []string{}
    for _,e := range exits {
		str = append(str, fmt.Sprintf("%v(%v)",e.Text(),e.Id()))
    }

    return "Exits: " + strings.Join(str,", ")
}

func Look(player *model.Player, params ...[]byte) string {
    fmt.Println("Look")
    if len(params) <= 1 {
        print("Look params 0")
        room := model.RoomById(player.RoomId())
        player.SendMsg(fmt.Sprintf("%v(%v)\n%v\n%v",room.Name(),room.Id(),room.Text(),LookExits(player)))
        return fmt.Sprintf("%v(%v)\n%v\n%v",room.Name(),room.Id(),room.Text(),LookExits(player))
    }

    fmt.Println("lookParams",params)
    switch string(params[1]) {
        case "exits":
            exits := model.GroupExitsBy("From")[player.RoomId()]
            str := []string{}
            for _,e := range exits {
			str = append(str, fmt.Sprintf("%v(%v)",e.Text(),e.Id()))
            }

            return "Exits: " + strings.Join(str,", ")
           
        default:
            return "Default!" 
    }
    return "No return"
}
