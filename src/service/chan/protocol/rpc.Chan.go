package protocol

import ()

type GameListRequest struct {
	Get bool `json:"get"`
}

type GameListResponse struct {
	List []Game `json:"l"`
}

type NameRequest struct {
	Get bool `json:"get"`
}

type NameResponse struct {
	Name string `json:"name"`
}
