package cmd

import "bytes"
import "../model"

type cmd map[string]func(*model.Player,...[]byte)
//var commands map[string]func(*model.Player,...[]byte) = map[string]func(*model.Player,...[]byte){"dig":Dig,"look":Look,"move":Move}
var commands cmd = cmd{"dig":Dig,"look":Look,"move":Move}
var login_commands cmd = cmd {"create":Create,"login":Login}

func ProcessCommand(player *model.Player, message []byte) {
   split := bytes.Split(message, []byte("`"))
   allparams := [][]byte{} //[]byte is a string, thus [][]byte is an array of strings(i.e. our parameters)
   for i, param := range split {
       if i % 2 == 0 {
           if trimmed := bytes.TrimSpace(param); len(trimmed) > 0 {
               splitted := bytes.Split(trimmed, []byte(" "))
               allparams = append(allparams, splitted...)
           }
       } else {
           allparams = append(allparams, param)
       }
   }

   if player.Base == nil {
       if command,ok := login_commands[string(allparams[0])]; ok {
           command(player, allparams...)
           return
       } else {
           player.SendMsg("Invalid command!")
           return
       }
      player.SendMsg("You are not logged in!")
      return
   }

   if command,ok := commands[string(allparams[0])]; ok {
       command(player, allparams...)
   } else {
       player.SendMsg("Invalid command!")
       return
   }

   //Can refactor the below into a new map[[]byte]func
   //if(bytes.Equal(allparams[0],[]byte("dig"))) {
   //    Dig(player,allparams...)
   //}
   //if(bytes.Equal(allparams[0],[]byte("look"))) {
   //    Look(player,allparams...)
   //}
   //if(bytes.Equal(allparams[0],[]byte("move"))) {
   //    Move(player,allparams...)
   //}
}
