package http_clients

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/amandavmanduca/fullcycle-gcr/interfaces"
	"github.com/amandavmanduca/fullcycle-gcr/structs"
)

type viaCepApiClient struct {
	httpClient interfaces.HttpClientInterface
}

func NewViaCepApiClient(httpClient interfaces.HttpClientInterface) interfaces.ViaCepApiInterface {
	return &viaCepApiClient{
		httpClient: httpClient,
	}
}

func (c *viaCepApiClient) GetAddress(ctx context.Context, cep string) (*structs.AddressResponse, error) {
	resp, err := c.httpClient.Get(ctx, fmt.Sprintf("/%s/json", cep), nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response structs.ViaCepAddressResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		return response.ToAddressResponse(), nil
	}
	return nil, errors.New("ERROR_GETTING_ADDRESS")
}
