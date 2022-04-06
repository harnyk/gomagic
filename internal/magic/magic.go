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

func NewMagicMap[V any]() *MagicMap[V] {
	return &MagicMap[V]{Value: nil}
}

func (m *MagicMap[V]) Set(key string, value V) *MagicMap[V] {
	if m.Value == nil {
		m.Value = make(map[string]V)
	}

	m.Value[key] = value
	return m
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
