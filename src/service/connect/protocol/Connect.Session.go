package protocol

type GetSessionRequest struct {
	Id uint64 `json:i`
}

type GetSessionResponse struct {
	Data []byte `json:d`
}

type SetSessionRequest struct {
	Id   uint64 `json:u`
	Data []byte `json:d`
}

type SetSessionResponse int
