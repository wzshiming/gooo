package protocol

import ()

type RegisterRequest struct {
	Username string `json:"u"`
	Password string `json:"p"`
	Email    string `json:"e"`
}

type ChangePwdRequest struct {
	Username    string `json:"u"`
	Password    string `json:"p"`
	NewPassword string `json:"n"`
}

type LogInRequest struct {
	Username string `json:"u"`
	Password string `json:"p"`
}

type LogOutRequest struct {
	LogOut bool `json:"l"`
}
