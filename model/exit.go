package model

import (
    "fmt"
    "../db"
    "github.com/go-pg/pg"
    "github.com/go-pg/pg/orm"
    "errors"
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
var exitsGroupedBy map[string]map[int]Exits

func (exit *Exit_) Insert() {
      if err := db.Db.Insert(exit); err != nil {
      panic(err)
    }
}

func (exit *Exit) Id() int {
    return exit.Base.Id
}

func (exit *Exit) To(params ...int) int {
    if len(params) > 0 {
        exit.Base.To = params[0]
    }
    return exit.Base.To
}

func (exit *Exit) From(params ...int) int {
    if len(params) > 0 {
        exit.Base.From = params[0]
    }
    return exit.Base.From
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
    exitsGroupedBy["From"][base.From] = append(exitsGroupedBy["From"][base.From],newExit)
    return newExit
}

func GroupExitsBy(prop string) map[int]Exits {
    switch prop {
        case "From":
            return exitsGroupedBy["From"]
        default:
            panic("Property not found!")
            //return allExits
    }
}

type myfunc func(*Exit) bool
func (exits *Exits) Find(f myfunc) (*Exit, error) {
    for _,e := range *exits {
        fmt.Println("bar")
        if f(e) == true {
            return e, nil
        }
    }
    return nil, errors.New("value_not_found")
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
    for i,_ := range allExits_ {
        
        newExit := &Exit{Base: &allExits_[i], State:make(map[string]interface{})}
        allExits[i] = newExit
        if _, ok := groupedbyfrom[newExit.From()]; !ok { 
            groupedbyfrom[newExit.From()] = Exits{newExit}
        } else {
            groupedbyfrom[newExit.From()] = append(groupedbyfrom[newExit.From()],newExit)
        }
    }
    //print("groupedbyfrom",groupedbyfrom[1][0].To())
    exitsGroupedBy = make(map[string]map[int]Exits)
    exitsGroupedBy["From"] = groupedbyfrom
    fmt.Println(allExits)
    
}
