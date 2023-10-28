package helpers

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonReMarshal(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonReMarshal(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonReMarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonReMarshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitArray(t *testing.T) {

	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{
			name:  "ok",
			input: []string{"UAH", "USD", "GBP"},
			want:  "UAH,USD,GBP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitArray(tt.input); got != tt.want {
				t.Errorf("SplitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidMailAddress(t *testing.T) {

	tests := []struct {
		name            string
		address         string
		expectedAddress string
		expectedResult  bool
	}{
		{
			name:           "ok",
			address:        "a@a.com",
			expectedResult: true,
		},
		{
			name:           "not ok",
			address:        "qwerxqwe",
			expectedResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ValidMailAddress(tt.address)
			assert.Equal(t, got, tt.address)
			assert.Equal(t, got1, tt.expectedResult)

		})
	}
}
