package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Ryan-18-system/desafio_multithreading_go_expert/internal/dto"
)

func GetCepBrasilApi(cep string) (*dto.Address, error) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v2/%s", cep)
	responseByte, err := executeRequest(url)
	if err != nil {
		return nil, err
	}
	address, err := parseJsonResponse(responseByte)
	if err != nil {
		return nil, err
	}
	return address, nil
}

func GetCepViaCeplApi(cep string) (*dto.Address, error) {

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	responseByte, err := executeRequest(url)
	if err != nil {
		return nil, err
	}
	address, err := parseJsonResponse(responseByte)
	if err != nil {
		return nil, err
	}
	return address, nil
}
func parseJsonResponse(jsonResponse []byte) (*dto.Address, error) {
	var address *dto.Address
	err := json.Unmarshal(jsonResponse, &address)
	if err != nil {
		return nil, err
	}
	return address, nil
}
func executeRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}
	responseByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// simulação para da timeout
	// time.Sleep(3 * time.Second)
	return responseByte, nil
}
