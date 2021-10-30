# Weather API Client

## Description
API Client with api key https://weatherapi.com/

## Install
```
go get github.com/ayd-ar/weather
```
## Example
```go
package main

import (
	"fmt"
	"log"

	"github.com/ayd-ar/weather"
)

func main() {
	wc := weather.NewClient("91ca83590f344dbe89a183734213010")

	vld, err := wc.GetWeather("Vladikavkaz")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(vld.Weather())
}
```