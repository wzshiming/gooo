package main

import (
	"rego/misc"
)

type Room struct {
	room *misc.Rooms
}

func NewRoom() *Room {
	r := Room{
		room: misc.NewRooms("Rooms"),
	}
	return &r
}
