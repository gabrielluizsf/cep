package cep

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name string
		cep  string
		want bool
	}{
		{
			name: "CEP válido #1",
			cep:  "01001-000",
			want: true,
		},
		{
			name: "CEP válido #2",
			cep:  "01001000",
			want: true,
		},
		{
			name: "CEP inválido #1",
			cep:  "01001-00",
			want: false,
		},
		{
			name:  "CEP inválido #2",
			cep:  "01001-0000",
			want: false,
		},
		{
			name: "CEP inválido #3",
			cep:  "01001-000a",
			want: false,
		},
		{
			name: "CEP inválido #4",	
			cep:  "01001-000-00",
			want: false,
		},
		{
			name: "CEP inválido #5",
			cep:  "01001-00000",
			want: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got := Validate(testCase.cep)
			assert.Equal(t, got, testCase.want)
		})
	}
}
