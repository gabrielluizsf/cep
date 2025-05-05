package api

import (
	"context"
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestFind(t *testing.T) {
	cep := "01001-000"
	data, err := Find(context.Background(), cep)
	if err != nil {
		t.Skipf("Erro ao consultar o ViaCEP: %s", err.Error())
	}
	assert.Equal(t, data.ZipCode, cep)
	assert.Equal(t, data.Street, "Praça da Sé")
	assert.Equal(t, data.Complement, "lado ímpar")
	assert.Equal(t, data.District, "Sé")
	assert.Equal(t, data.City, "São Paulo")
	assert.Equal(t, data.State, "SP")
	assert.Equal(t, data.IBGE, "3550308")
	assert.Equal(t, data.GIA, "1004")
	assert.Equal(t, data.DDD, "11")
	assert.Equal(t, data.SIAFI, "7107")
	assert.Equal(t, data.Region, "Sudeste")
}
