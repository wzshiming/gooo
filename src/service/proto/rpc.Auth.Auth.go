package proto

import ()

type UnregisterRequest struct {
	Username string `json:"u"`
	Password string `json:"p"`
}

type RegisterRequest struct {
	Username string `json:"u"`
	Password string `json:"p"`
	Email    string `json:"e"`
}

type LogInRequest struct {
	Username string `json:"u"`
	Password string `json:"p"`
}

type LogOutRequest struct {
	LogOut bool `json:"l"`
}

type ChangePwdRequest struct {
	Username    string `json:"u"`
	Password    string `json:"p"`
	NewPassword string `json:"n"`
}
