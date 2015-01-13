package protocol

import (
//"time"
)

type VersionsRequest struct {
	Version string `json:"version"`
}

type VersionsResponse struct {
	Versions []string `json:"versions"`
}
