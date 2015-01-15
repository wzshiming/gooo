package protocol

import ()

type RegisterRequest struct {
	Email    string `json:"e"`
	Username string `json:"u"`
	Password string `json:"p"`
}

type OkResponse struct {
	Ok bool `json:"ok"`
}

type ChangePwdRequest struct {
	Username    string `json:"u"`
	Password    string `json:"p"`
	NewPassword string `json:"n"`
}
