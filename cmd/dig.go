package cmd

import "../model"

func Dig(player *model.Player, params ...[]byte) {
    name := params[1]
    text := params[2]
    //state := map[string]interface{}{}

    room := model.Room_{
        Name: string(name),
        Text: string(text),
     //   State: &state,
    }

    model.CreateRoom(&room)
    //room.Insert()
}
