package magic

import "encoding/json"

//---------------------------

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

//---------------------------

//MagicMap unmarshals itself to a structure with a Value field of type V.
//In case of an error, the Value field is set to nil.
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

//---------------------------

//BareMap unmarshals itself to an empty map in case of an error.
type BareMap[V any] map[string]V

func (m *BareMap[V]) UnmarshalJSON(data []byte) error {
	var value map[string]V

	err := json.Unmarshal(data, &value)
	if err != nil {
		// m = nil
		return nil
	}
	*m = value
	return nil
}

// //Creates a new PracticalMap[V]
// func NewPracticalMap[V any]() *PracticalMap[V] {
// 	return &PracticalMap[V]{}
// }

// func (m *PracticalMap[V]) Set(key string, value V) *PracticalMap[V] {
// 	if m == nil {
// 		m = NewPracticalMap[V]()
// 	}

// 	(*m)[key] = value
// 	return m
// }
