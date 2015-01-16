package protocol

type GetOnlineRequest struct {
	UserId int64 `json:u`
}

type GetOnlineResponse struct {
	Online bool `json:o`
}

type SetOnlineRequest struct {
	UserId int64 `json:u`
	Online bool  `json:o`
}

type SetOnlineResponse int
