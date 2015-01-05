package coding

import (
	"coding/bson"
	"encoding/json"
)

// encoding json and bson
var (
	Unmarshal = json.Unmarshal
	Marshal   = json.Marshal
)

func SetCodingBson(){
	Unmarshal = bson.Unmarshal
	Marshal   = bson.Marshal
}

func SetCodingJson(){
	Unmarshal = json.Unmarshal
	Marshal   = json.Marshal
}
