package protocol

import (
	//"time"
	"gooo/configs"
)

type InitRequest struct {
	Conf  configs.Configs `json: conf`
	State int             `json: state`
}
