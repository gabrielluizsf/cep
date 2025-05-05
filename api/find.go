package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/i9si-sistemas/nine"
	"github.com/i9si-sistemas/nine/pkg/client"

	"github.com/i9si-sistemas/cep"
)

// CEP represents the structure of the data returned by the ViaCEP API
type CEP struct {
	ZipCode    string `json:"cep"`
	Street     string `json:"logradouro"`
	Complement string `json:"complemento"`
	District   string `json:"bairro"`
	City       string `json:"localidade"`
	State      string `json:"uf"`
	IBGE       string `json:"ibge"`
	GIA        string `json:"gia"`
	DDD        string `json:"ddd"`
	Region     string `json:"regiao"`
	SIAFI      string `json:"siafi"`
}

// Find queries ViaCEP and returns zip code data
func Find(ctx context.Context, zipCode string) (*CEP, error) {
	if !cep.Validate(zipCode) {
		return nil, fmt.Errorf("formato de CEP inválido")
	}

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipCode)

	res, err := nine.New(ctx).Get(url, &client.Options{})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro ao consultar o ViaCEP. Status: %d", res.StatusCode)
	}

	var data CEP
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
