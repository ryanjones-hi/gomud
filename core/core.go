package core

import "fmt"

type EventParams struct {
    FireOnce bool
}

type EventContext struct {
    previousActionType ActionType
    abort bool
}

type Action *func(*EventContext)
type Event *func(*EventContext)
type EventS []Event

//Example: Move, Attack, etc.
type ActionType string

//Example: model.Room2255, model.Player83
type ObjectGuid string

const (
    NOMINATIVE = iota
    ACCUSATIVE
    DATIVE
    LOCATIVE
    INSTRUMENTATIVE
    ABLATIVE
    ALLATIVE
    GENITIVE
    COUNT
    VALUE
    VAR1
    VAR2
    VAR3
)
type ParamType int
type ActionParams map[ParamType]ObjectGuid

var events map[ActionType]map[ParamType]map[ObjectGuid]EventS

func TriggerEvents(ctx *EventContext, params ActionParams, at ActionType) {
    //Now we need to check nominative, accusative, dative, and locative one by one to check if there are any events to trigger.
    for k,v := range params {
        TriggerEvent(ctx,at,k,v)
    }
}

func TriggerEvent(ctx *EventContext, at ActionType, pt ParamType, guid ObjectGuid) {
    if _, ok := events[at]; ok {
    if _, ok := events[at][pt]; ok {
    if eventlist, ok := events[at][pt][guid]; ok {
        for _, e := range eventlist {
            (*e)(ctx)
            //fmt.Println(e)
        }
    }}}
}

type ActionBuilder func(ActionParams)(Action,ActionType)

//Could also specify additional conditions under which this event will fire
//eparams *EventParams
func BuildEvent(actionbuilder ActionBuilder, params ActionParams) (Event) {

    action, actiontype := actionbuilder(params)
    e := func(ctx *EventContext) {
        //before action. e.g. (trigger "before" EventS)
        TriggerEvents(ctx, params,"before" + actiontype)
        (*action)(ctx)
        //after action
        fmt.Println(actiontype)
        TriggerEvents(ctx, params, actiontype)
    }
    return &e
}

func Atk(params ActionParams) (Action,ActionType) {
    at := "Atk"
    a := func(ctx *EventContext){
        fmt.Println("Atk action fired!")
        fmt.Println(params[ACCUSATIVE])
        //fmt.Println(params[NOMINATIVE])
        //fmt.Println(params[ACCUSATIVE])
    }
    return &a, ActionType(at)
    //return BuildActionHandler(&f, params)
}

func RegisterEvent(e Event, at ActionType, pt ParamType, o ObjectGuid) {
    if _, ok := events[at]; !ok {
        events[at] = make(map[ParamType]map[ObjectGuid]EventS)
        events[at][pt] = make(map[ObjectGuid]EventS)
        events[at][pt][o] = EventS{e}
    } else if _, ok := events[at][pt]; !ok {
        events[at][pt] = make(map[ObjectGuid]EventS)
        events[at][pt][o] = EventS{e}
    } else if _, ok := events[at][pt][o]; !ok {
        events[at][pt][o] = EventS{e}
    } else {
        events[at][pt][o] = append(events[at][pt][o],e)
    }
}

func Init() {
    events = make(map[ActionType]map[ParamType]map[ObjectGuid]EventS)
}

func main() {
    Init()
    params := ActionParams{ACCUSATIVE:"bob"}
    e := BuildEvent(Atk, params)
    //Now let us begin: Time to try to register this thing.

    RegisterEvent(e, "Atk", ACCUSATIVE, "bob")
    RegisterEvent(e, "Atk", ACCUSATIVE, "jim")
    ctx := EventContext{"Atk",true}
    (*e)(&ctx)
    //func1 := AtkF(&ActionParams{Nominative:"foo",Accusative:"bar"})
    //func2 := AtkF(&ActionParams{Nominative:"foo2",Accusative:"bar2"})
    //locative_EventS["room1"] = map[ActionType]EventS{}
    //locative_EventS["room1"]["atk"] = EventS{func1, func2}
    //ctx := EventContext{"run",false}
    //for _,Action := range locative_EventS["room1"]["atk"] {
    //    (*Action)(&ctx)
    //}
}
