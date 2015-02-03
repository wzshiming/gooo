package protocol

import ()

type LeaveRequest int

type LeaveResponse int

type InterruptRequest struct {
	RoomId int
	SeatId int
}

type InterruptResponse int
