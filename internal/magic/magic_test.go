package magic

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMagic_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name     string
		expected int64
		args     args
		wantErr  bool
	}{
		{
			name:     "UnmarshalJSON",
			expected: 1,
			args: args{
				data: []byte("1"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Magic{}
			if err := json.Unmarshal(tt.args.data, m); (err != nil) != tt.wantErr {
				t.Errorf("Magic.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if m.Value != tt.expected {
				t.Errorf("Magic.UnmarshalJSON() = %v, want %v", m.Value, tt.expected)
			}
		})
	}
}

func TestMagicBoolMap_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name     string
		expected *MagicMap[bool]
		args     args
		wantErr  bool
	}{
		{
			name:     "UnmarshalJSON",
			expected: NewMagicMap[bool]().Set("a", true),
			args: args{
				data: []byte(`{"a":true}`),
			},
		},
		{
			name:     "UnmarshalJSON",
			expected: NewMagicMap[bool](),
			args: args{
				data: []byte(`[]`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MagicMap[bool]{}
			if err := json.Unmarshal(tt.args.data, m); (err != nil) != tt.wantErr {
				t.Errorf("MagicBoolMap.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(m, tt.expected) {
				t.Errorf("MagicBoolMap.UnmarshalJSON() = %v, want %v", m, tt.expected)
			}
		})
	}
}

func TestMagicBoolMap_UnmarshalJSON_2(t *testing.T) {
	type Attributes struct {
		Features MagicMap[MagicMap[bool]]
	}

	type args struct {
		data []byte
	}
	tests := []struct {
		name     string
		expected Attributes
		args     args
		wantErr  bool
	}{
		{
			name: "UnmarshalJSON",
			expected: Attributes{
				Features: *NewMagicMap[MagicMap[bool]]().
					Set("a", *NewMagicMap[bool]().
						Set("b", true)),
			},
			args: args{
				data: []byte(`{
					"features": {
						"a": {
							"b":true
						}
					}
				}`),
			},
		},
		{
			name: "UnmarshalJSON: wrong map type",
			expected: Attributes{
				Features: *NewMagicMap[MagicMap[bool]]().
					Set("a", *NewMagicMap[bool]()),
			},
			args: args{
				data: []byte(`{
					"features": {
						"a": []
					}
				}`),
			},
		}, {
			name: "UnmarshalJSON: wrong map parent type",
			expected: Attributes{
				Features: *NewMagicMap[MagicMap[bool]](),
			},
			args: args{
				data: []byte(`{
					"features": []
				}`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Attributes{}
			if err := json.Unmarshal(tt.args.data, &m); (err != nil) != tt.wantErr {
				t.Errorf("MagicBoolMap.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(m, tt.expected) {
				t.Errorf("MagicBoolMap.UnmarshalJSON() = %v, want %v", m, tt.expected)
			}
		})
	}
}
