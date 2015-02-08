package protocol

import ()

type LeaveRequest int

type LeaveResponse int

type InterruptLeaveRequest struct {
	RoomId int
	SeatId int
}

type InterruptLeaveResponse int
