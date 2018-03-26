package model

type Name string
type Description string

type EntityGuid string
type EntityType string
type EntityId int

//This is the flagship interface for all of our models. This wil allow us to set up events for these.
type IEntity interface {
    guid() EntityGuid
    type_() EntityType
    id() EntityId
}
type IEntities []IEntity

type IContainer interface {
    contents() IEntities
}

//A possession is an item that can be carried by players/mobs
type IPossession interface {
//    weight() int
//    volume() int
    posessor() IContainer
    name() Name
    desc() Description
}
type Inventory []IPossession

type ILocation interface {
    name() Name
    desc() Description
    mobiles() IMobiles
    contents() IEntities
    exits() IExits
}
type ILocations []ILocation

type IExit interface {
    name() Name
    desc() Description
    destination() ILocation
    location() ILocation
}
type IExits []IExit

type IPersistent interface {
    save() bool
}

type IImmobile interface {
    location() ILocation
    desc() Description
    name() Name
    obtainable() bool
}
type IImobiles []IImmobile

type IMobile interface {
    location() ILocation
    desc() Description
    name() Name
    inventory() Inventory
}
type IMobiles []IMobile
