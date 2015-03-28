package agent

import (
	"rego"

	//"net/http"
)

type Request struct {
	Request *rego.EncodeBytes
	Session *Session
	Head    []byte
}

type Response struct {
	Error    string
	Head     []byte
	Data     *rego.EncodeBytes
	Coverage *rego.EncodeBytes
	Response *rego.EncodeBytes
}

func (re Response) Hand(user *User, head []byte) error {
	var ret []byte
	if re.Coverage != nil {
		user.Data = re.Coverage
	} else if re.Data != nil {
		user.Data = rego.SumJson(user.Data, re.Data)
	}
	if re.Error != "" {
		ret = []byte(`{"error":"` + re.Error + `"}`)
	} else if re.Response != nil {
		ret = re.Response.Bytes()
	} else {
		ret = []byte(`{"error":""}`)
	}
	if re.Head != nil && len(re.Head) != 0 {
		ret = append(re.Head, ret...)
	} else {
		ret = append(head, ret...)
	}
	return user.WriteMsg(ret)
}
