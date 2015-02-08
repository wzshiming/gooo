package protocol

type SendRequest struct {
	Clients []uint64 `json:c`
	Data    []byte   `json:d`
}

type SendResponse int
