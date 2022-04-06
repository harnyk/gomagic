package magic

import "encoding/json"

type Magic struct {
	Value int64
}

func (m *Magic) UnmarshalJSON(data []byte) error {
	var value int64
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}
	m.Value = value
	return nil
}

type MagicMap[V any] struct {
	Value map[string]V
}

func (m *MagicMap[V]) UnmarshalJSON(data []byte) error {
	var value map[string]V
	err := json.Unmarshal(data, &value)
	if err != nil {
		return nil
	}
	m.Value = value
	return nil
}
