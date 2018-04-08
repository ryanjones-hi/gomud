package cmd

import "../model"
import "fmt"

func Look(player *model.Player, params ...[]byte) string {
    fmt.Println(player.Room().Text())
    return player.Room().Name() + "\n" + player.Room().Text()
}
