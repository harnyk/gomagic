package magic

import (
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
			if err := m.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
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
			expected: &MagicMap[bool]{Value: map[string]bool{"a": true}},
			args: args{
				data: []byte(`{"a":true}`),
			},
		},
		{
			name:     "UnmarshalJSON",
			expected: &MagicMap[bool]{Value: nil},
			args: args{
				data: []byte(`[]`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MagicMap[bool]{}
			if err := m.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("MagicBoolMap.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(m, tt.expected) {
				t.Errorf("MagicBoolMap.UnmarshalJSON() = %v, want %v", m, tt.expected)
			}
		})
	}
}
