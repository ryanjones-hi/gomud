package cmd

import "../model"
import "fmt"
import "strings"

func Look(player *model.Player, params ...[]byte) string {
    if len(params) == 0 {
        return player.Room().Name() + "\n" + player.Room().Text() 
    }

    fmt.Println("lookParams",params)
    switch string(params[1]) {
        case "exits":
            //fmt.Println("groupedExits",model.GroupExitsBy("From")[player.Room().Id()][0].To())
            exits := model.GroupExitsBy("From")[player.Room().Id()]
            str := []string{}
            for _,e := range exits {
                str = append(str, e.Text())
            }
            fmt.Println("built string",strings.Join(str,", "))

            return "Exits: " + strings.Join(str,", ")
            //return player.Room().Name() + '\n' + player.Room().Text() 
            //return model.GroupExitsBy("From")[player.Room().Id()]
           
        default:
            return "Default!" 
    }

        
    fmt.Println(player.Room().Text())
    return "No return"
}
