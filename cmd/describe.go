package cmd

import "../model"
import "strconv"
import "fmt"

func DescribeRoom(player *model.Player, params ...[]byte) {
    if len(params) < 2 {
        player.SendMsg("Format: DESCRIBEROOM description") 
        return
    }
    room, err := model.AllRooms.Find(func(_room *model.Room)bool { return _room.Id() == player.RoomId() })
    if err != nil {
        player.SendMsg("Wasn't able to redesc the room")
        return
    }
   
    room.Desc(string(params[1]))
    player.SendMsg("Room description has been updated!")
}

func DescribeExit(player *model.Player, params ...[]byte) {
    if len(params) < 3 {
        player.SendMsg("Format: DESCRIBEEXIT (entity id) description") 
        return
    }

    exitId, err := strconv.Atoi(string(params[1]))
    if err != nil {
        player.SendMsg("Id must be a number")
        return
    }
    fmt.Println("DescribeExit:",exitId)
    exit, err := model.AllExits.Find(func(_exit *model.Exit)bool { return _exit.Id() == exitId })

    if err != nil {
        player.SendMsg("Exit with that id was not found")
        return
    }

    exit.Desc(string(params[2]))
    player.SendMsg("Exit description has been updated!")
}

func Describe(player *model.Player, params ...[]byte) {
    if len(params) < 3 {
        player.SendMsg("Format: DESCRIBE (me/entity id) description") 
        return
    }
    
    if string(params[1]) == "me" {
        player.Desc(string(params[2]))
        player.SendMsg("Description changed")
    }
}
