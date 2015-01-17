package handel

import (
	"fmt"
)

func ErrorInfo(s string) []byte {
	return []byte(fmt.Sprintf("{\"e\":\"%s\"}", s))
}

func OkInfo() []byte {
	return []byte("{\"e\":\"\"}")
}
