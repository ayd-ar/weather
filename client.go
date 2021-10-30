package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (r Client) GetWeather(city string) (response, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", r.apiKey, city)

	resp, err := r.httpClient.Get(url)
	if err != nil {
		return response{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response{}, err
	}

	var responseData response
	if err := json.Unmarshal(body, &responseData); err != nil {
		return response{}, err
	}

	return responseData, nil
}
