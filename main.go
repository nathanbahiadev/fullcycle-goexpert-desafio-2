package main

import (
	"context"
	"errors"
	"fmt"
	"fullcycle_multithreading/services"
	"log"
	"time"
)

func main() {
	var cep services.CEP

	fmt.Print("Informe o CEP que deseja consultar: ")
	fmt.Scan(&cep)

	if err := cep.Validate(); err != nil {
		log.Fatal(err)
	}

	cep.Clear()
	address, err := GetAddress(cep, services.BrasilAPIService)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("============== Resultados encontrados! ==============")
	fmt.Println("Cep:", address.GetCep())
	fmt.Println("Street:", address.GetStreet())
	fmt.Println("Neighborhood:", address.GetNeighborhood())
	fmt.Println("City:", address.GetCity())
	fmt.Println("State:", address.GetState())
	fmt.Println("Service:", address.GetService())
	fmt.Println("=====================================================")
}

func GetAddress(cep services.CEP, funcs ...services.TGetAddressFuncService) (services.TAddress, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second))
	defer cancel()

	addressCh := make(chan services.TAddress)

	for _, service := range funcs {
		go service(ctx, addressCh, cep)
	}

	select {
	case address := <-addressCh:
		return address, nil

	case <-ctx.Done():
		return nil, errors.New("timeout ao consultar CEP")
	}
}
