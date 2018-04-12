package cmd

import "../model"

func Dig(player *model.Player, params ...[]byte) {
    currentRoomId := player.RoomId()
    name := params[1]
    text := params[2]
    //state := map[string]interface{}{}

    room := model.Room_{
        Name: string(name),
        Text: string(text),
     //   State: &state,
    }

    nextRoom := model.CreateRoom(&room)
    nextRoomId := nextRoom.Id()


    model.CreateExit(&model.Exit_{To:nextRoomId,From:currentRoomId,Text:string(text)})
    model.CreateExit(&model.Exit_{To:currentRoomId,From:nextRoomId,Text:model.RoomById(player.RoomId()).Name()})
    //make exits to and from the new room
    //room.Insert()
}
