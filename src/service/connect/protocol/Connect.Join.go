package protocol

type JoinRequest struct {
	Id uint64 `json:"id"`
}

type JoinResponse struct {
	Sum      int    `json:"sum"`
	Response []byte `json:"response"`
}
