package protocol

type GetSessionRequest struct {
	Id uint `json:i`
}

type GetSessionResponse struct {
	Data []byte `json:d`
}

type SetSessionRequest struct {
	Id   uint   `json:u`
	Data []byte `json:d`
}

type SetSessionResponse int
