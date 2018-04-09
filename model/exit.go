package model

import (
    "fmt"
    "../db"
    "github.com/go-pg/pg"
    "github.com/go-pg/pg/orm"
)

type Exit_ struct {
    Id int
    From int
    To int
    Text string
}

type Exit struct {
    Base *Exit_
    State map[string]interface{}
}
type Exits []*Exit
var allExits Exits
var groupedBy map[string]map[int]Exits

func (exit *Exit_) Insert() {
      if err := db.Db.Insert(exit); err != nil {
      panic(err)
    }
}

func (exit *Exit) To(params ...int) int {
    if len(params) > 0 {
        exit.Base.To = params[0]
    }
    return exit.Base.To
}

func (exit *Exit) Text(params ...string) string {
    if len(params) > 0 {
        exit.Base.Text = params[0]
    }
    return exit.Base.Text
}

func CreateExit(base *Exit_) *Exit {
    base.Insert()
    newExit := &Exit{Base: base, State:make(map[string]interface{})}
    groupedBy["From"][base.From] = append(groupedBy["From"][base.From],newExit)
    return newExit
}

func GroupExitsBy(prop string) map[int]Exits {
    switch prop {
        case "From":
            return groupedBy["From"]
        default:
            panic("Property not found!")
            //return allExits
    }
}

func init() {
    Db := pg.Connect(&pg.Options{
        User:      "gomud",
        Password:  "gomud",
        Database:  "gomud",
    })
    Db.CreateTable((*Exit_)(nil),&orm.CreateTableOptions{})
    var allExits_ []Exit_
    err := Db.Model(&allExits_).Select()
    if err!= nil {
       panic(err)
    }
    allExits = make(Exits, len(allExits_))

    groupedbyfrom := make(map[int]Exits)
    for i,base := range allExits_ {
        
        newExit := &Exit{Base: &base, State:make(map[string]interface{})}
        allExits[i] = newExit
        if exits, ok := groupedbyfrom[base.From]; !ok { 
            groupedbyfrom[base.From] = Exits{newExit}
        } else {
            exits = append(exits,newExit)
        }
    }
    //print("groupedbyfrom",groupedbyfrom[1][0].To())
    groupedBy = make(map[string]map[int]Exits)
    groupedBy["From"] = groupedbyfrom
    fmt.Println(allExits)
    
}

