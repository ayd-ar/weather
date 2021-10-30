package weather

import "fmt"

type response struct {
	location `json:"location"`
	current  `json:"current"`
}

type location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type current struct {
	TempC float64 `json:"temp_c"`
}

func (r response) Weather() string {
	return fmt.Sprintf("%.0f°C - %s, %s", r.TempC, r.Name, r.Country)
}
