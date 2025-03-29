package main

import (
	"fmt"
	"time"

	"github.com/Ryan-18-system/desafio_multithreading_go_expert/internal/service"
)

func main() {

	channelBrasilApi := make(chan string)
	channelViaCepApi := make(chan string)
	go func() {

		result, err := service.GetCepViaCeplApi("58025140")
		if err == nil {
			channelViaCepApi <- fmt.Sprintf("Get ViaCep api %s", result.FormatedAddressViaCepApi())
		}
		close(channelViaCepApi)
	}()
	go func() {

		result, err := service.GetCepBrasilApi("58041150")
		if err == nil {
			channelBrasilApi <- fmt.Sprintf("Get Brasil api %s", result.FormatedAddressBrasilApi())
		}
		close(channelBrasilApi)
	}()

	select {
	case result := <-channelBrasilApi:
		fmt.Println(result)
	case result := <-channelViaCepApi:
		fmt.Println(result)
	case <-time.After(1 * time.Second):
		println("timeout")
	}
}
