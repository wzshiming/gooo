package main

import (
	"rego/misc"
)

type Room struct {
	room *misc.Rooms
}

func NewRoom(size int) *Room {
	r := Room{
		room: misc.NewRooms("Rooms", size),
	}
	return &r
}
