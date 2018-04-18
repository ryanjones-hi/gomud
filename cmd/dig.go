package cmd

import "../model"

func Dig(player *model.Player, params ...[]byte) {
    if len(params) < 3 {
        player.SendMsg("DIG Format: dig roomname roomdesc")
        return
    }
    currentRoomId := player.RoomId()
    name := params[1]
    text := params[2]
    //state := map[string]interface{}{}

    room := model.Room_{
        Name: string(name),
        Desc: string(text),
     //   State: &state,
    }

    nextRoom := model.CreateRoom(&room)
    nextRoomId := nextRoom.Id()


    model.CreateExit(&model.Exit_{To:nextRoomId,From:currentRoomId,Desc:string(text),Name:string(text)})
    model.CreateExit(&model.Exit_{To:currentRoomId,From:nextRoomId,Desc:model.RoomById(player.RoomId()).Name(),Name:string(text)})
    SysMove(player, nextRoomId)
    //make exits to and from the new room
    //room.Insert()
}
