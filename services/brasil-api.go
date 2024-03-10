package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type BrasilAPIResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func (c *BrasilAPIResponse) GetCep() string {
	return c.Cep
}

func (c *BrasilAPIResponse) GetState() string {
	return c.State
}

func (c *BrasilAPIResponse) GetCity() string {
	return c.City
}

func (c *BrasilAPIResponse) GetNeighborhood() string {
	return c.Neighborhood
}

func (c *BrasilAPIResponse) GetStreet() string {
	return c.Street
}

func (c *BrasilAPIResponse) GetService() string {
	return "BrasilAPI"
}

func BrasilAPIService(ctx context.Context, ch chan TAddress, cep CEP) {

	brasilAPIResponse := BrasilAPIResponse{}
	var address TAddress = &brasilAPIResponse

	req, err := http.NewRequestWithContext(ctx, "GET", "https://brasilapi.com.br/api/cep/v1/"+string(cep), nil)

	if err != nil {
		log.Println(err.Error())
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println(err.Error())
	}

	err = json.NewDecoder(res.Body).Decode(&brasilAPIResponse)

	if err != nil {
		log.Println(err.Error())
	}

	defer res.Body.Close()

	if address.GetCep() != "" {
		ch <- address
	}
}
