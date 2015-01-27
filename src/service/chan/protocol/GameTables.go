package protocol

import ()

type Game struct {
	Users   []uint64
	ZoneId  uint64
	Name    string
	Started bool
}
