package route

import (
	"gooo/configs"
	"gooo/encoder"
)

func ClientRequest(i1, i2, i3 uint8, arge interface{}) []byte {
	sms := []byte{0, i1, i2, i3}
	b, err := encoder.Encode(arge)
	if err != nil {
		return []byte{}
	}
	return append(sms, b...)
}

func ClientRequestForm(rc *configs.Configs, c1, c2, c3 string, arge interface{}) []byte {
	i1, i2, i3 := rc.Rc.FindIndex(c1, c2, c3)
	return ClientRequest(i1, i2, i3, arge)
}
