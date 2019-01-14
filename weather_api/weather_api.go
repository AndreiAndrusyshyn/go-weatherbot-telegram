package weather_api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ind struct {
	Coord struct {
		Lon float32
		Lat float32
	}

	Wheather struct {
		Main        string
		Description string
	}

	Main struct {
		Temp     float32
		Pressure float32
		Temp_min float32
		Temp_max float32
	}

	Wind struct {
		Speed float32
	}
	Id   int
	Name string
	Cod  int
}

func Get_weather(city string) (result ind) {

	part1 := "http://api.openweathermap.org/data/2.5/weather?q="
	part2 := city
	url := string(part1 + part2 + "&APPID=TOKEN")

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal("NewRequest:", err)

	}
	res, _ := spaceClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)
	ind1 := ind{}

	jsonErr := json.Unmarshal(body, &ind1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return ind1
}
