package cmd

import "../model"
import "strconv"

func RenameRoom(player *model.Player, params ...[]byte) {
    if len(params) < 2 {
        player.SendMsg("Format: RENAMEROOM newname")
        return
    }
    
    room, _ := model.AllRooms.Find(func(_room *model.Room)bool { return _room.Id() == player.RoomId() })

    room.Name(string(params[1]))
    player.SendMsg("Renamed room!")
        
}

func RenameExit(player *model.Player, params ...[]byte) {
    if len(params) < 3 {
        player.SendMsg("Format: RENAMEEXIT exitid newname")
        return
    }

    exitId, err := strconv.Atoi(string(params[1]))
    if err != nil {
        player.SendMsg("Exit could not be renamed! Parameter not id number?")
        return
    }
    
    exit, err := model.AllExits.Find(func(_exit *model.Exit)bool { return _exit.Id() == exitId })
    if err != nil {
        player.SendMsg("Couldn't find exit")
        return
    }

    exit.Name(string(params[2]))
    player.SendMsg("Renamed exit!")
        
}
