package state

import ()

const (
	LS_UNALLOW = uint8(iota)
	LS_GUEST   //未登入的
	LS_LOGIN   //登入了的
	LS_INGAME  //在游戏中
	LS_OFFLINE //游戏中断线
)

const (
	LF_NONE  = uint32(1 << iota)
	LF_COMM  = uint32(1 << iota)
	LF_VIP   = uint32(1 << iota)
	LF_ADMIN = uint32(1 << iota)
)
