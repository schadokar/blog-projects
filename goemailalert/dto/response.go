package dto

type WeatherAPIResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name string `json:"name"`
}

type Current struct {
	Temperature int `json:"temperature"`
}
