package coding

import (
	"errors"
	"coding/bson"
	
)
// RawMessage is a raw encoded JSON or BSON object.
// It implements Marshaler and Unmarshaler and can
// be used to delay JSON or BSON decoding or precompute a JSON or BSON encoding.
type RawMessage []byte

// MarshalJSON returns *m as the JSON encoding of m.
func (m *RawMessage) MarshalJSON() ([]byte, error) {
	return *m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawMessage) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("coding.RawMessage: UnmarshalJSON on nil pointer")
	}
	*m = data[:]
	return nil
}

// GetBSON returns *m as the BSON encoding of m.
func (m *RawMessage) GetBSON() (interface{}, error){
	return *m, nil
}

// SetBSON sets *m to a copy of data.
func (m *RawMessage) SetBSON(raw bson.Raw) error {
	if m == nil {
		return errors.New("coding.RawMessage: SetBSON on nil pointer")
	}
	*m = raw.Data[:]
	return nil
}


