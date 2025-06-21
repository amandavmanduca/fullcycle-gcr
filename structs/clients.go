package structs

type CepOrigin string

var (
	VIA_CEP CepOrigin = "Via Cep API"
)

type ViaCepAddressResponse struct {
	Cep          string `json:"cep"`
	Street       string `json:"logradouro"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
	Estado       string `json:"estado"`
}

func (a *ViaCepAddressResponse) ToAddressResponse() *AddressResponse {
	return &AddressResponse{
		Address: Address{
			Cep:          a.Cep,
			State:        a.State,
			City:         a.City,
			Neighborhood: a.Neighborhood,
			Street:       a.Street,
		},
		Origin: VIA_CEP,
	}
}

type Address struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

type AddressResponse struct {
	Address Address
	Origin  CepOrigin `json:"origin"`
}

type CurrentWeather struct {
	TempC float64 `json:"temp_c"`
}
type WeatherResponse struct {
	Current CurrentWeather `json:"current"`
}
