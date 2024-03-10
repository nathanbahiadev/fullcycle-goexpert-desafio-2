package services

import (
	"context"
	"errors"
	"regexp"
)

type CEP string

func (c *CEP) Validate() error {
	re, _ := regexp.Compile(`^[\d]{2}[.]?[\d]{3}[-]?[\d]{3}$`)
	if !re.Match([]byte(*c)) {
		return errors.New("informe um CEP no formato " + re.String())
	}

	return nil
}

func (c *CEP) Clear() {
	re, _ := regexp.Compile(`\D`)
	*c = CEP(re.ReplaceAll([]byte(*c), []byte("")))
}

type TAddress interface {
	GetCep() string
	GetState() string
	GetCity() string
	GetNeighborhood() string
	GetStreet() string
	GetService() string
}

type TGetAddressFuncService func(ctx context.Context, ch chan TAddress, cep CEP)
