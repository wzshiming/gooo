package gooo

import ()

type Methods struct {
	MethodsResponse
}

func NewMethods(allow uint32, unallow uint32, method ...string) *Methods {
	return &Methods{
		MethodsResponse{
			Allow:   allow,
			Unallow: unallow,
			Method:  method,
		},
	}
}

func (s *Methods) Method(args MethodsRequest, reply *MethodsResponse) error {
	*reply = s.MethodsResponse
	return nil
}
