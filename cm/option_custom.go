package cm

import (
	"bytes"
	"encoding/json"
)

var nullJSON = []byte("null")

func (o option[T]) MarshalJSON() ([]byte, error) {
	if !o.isSome {
		return nullJSON, nil
	}

	return json.Marshal(o.some)
}

func (o *option[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullJSON) {
		*o = option[T]{}
		return nil
	}

	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*o = option[T]{isSome: true, some: v}
	return nil
}
