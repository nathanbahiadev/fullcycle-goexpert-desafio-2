package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type ViaCepResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"uf"`
	City         string `json:"localidade"`
	Neighborhood string `json:"bairro"`
	Street       string `json:"logradouro"`
}

func (c *ViaCepResponse) GetCep() string {
	return c.Cep
}

func (c *ViaCepResponse) GetState() string {
	return c.State
}

func (c *ViaCepResponse) GetCity() string {
	return c.City
}

func (c *ViaCepResponse) GetNeighborhood() string {
	return c.Neighborhood
}

func (c *ViaCepResponse) GetStreet() string {
	return c.Street
}

func (c *ViaCepResponse) GetService() string {
	return "VIACep"
}

func ViaCepService(ctx context.Context, ch chan TAddress, cep CEP) {

	viaCepResponse := ViaCepResponse{}
	var address TAddress = &viaCepResponse

	req, err := http.NewRequestWithContext(ctx, "GET", "http://viacep.com.br/ws/"+string(cep)+"/json/", nil)

	if err != nil {
		log.Println(err.Error())
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println(err.Error())
	}

	err = json.NewDecoder(res.Body).Decode(&viaCepResponse)

	if err != nil {
		log.Println(err.Error())
	}

	defer res.Body.Close()

	if address.GetCep() != "" {
		ch <- address
	}
}
