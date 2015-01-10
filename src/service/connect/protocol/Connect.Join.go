package protocol

type JoinRequest struct {
	Id uint `json: id`
}

type JoinResponse struct {
	Sum      int    `json:sum`
	Response []byte `json:response`
}
