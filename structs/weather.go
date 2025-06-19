package structs

type Weather struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

func NewWeatherFromCelsius(tempC float64) *Weather {
	return &Weather{
		TempC: tempC,
		TempF: (tempC * 1.8) + 32,
		TempK: tempC + 273,
	}
}
