package main

import (
	"encoding/json"
	"os"
	"strings"
)

type Airport struct {
	ICAO      string  `json:"icao"`
	IATA      string  `json:"iata"`
	Name      string  `json:"name"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	Country   string  `json:"country"`
	Elevation int     `json:"elevation"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	TZ        string  `json:"tz"`
}

var cityMap map[string][]Airport

func createCityMap(filePath string) (map[string][]Airport, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var airports map[string]Airport
	err = json.Unmarshal(data, &airports)
	if err != nil {
		return nil, err
	}

	cityMap := make(map[string][]Airport)
	for _, airport := range airports {
		city := strings.ToLower(airport.City)
		cityMap[city] = append(cityMap[city], airport)
	}

	return cityMap, nil
}

func searchCities(cityMap map[string][]Airport, searchTerm string) []Airport {
	searchTerm = strings.ToLower(searchTerm)
	var results []Airport

	for city, airports := range cityMap {
		if strings.Contains(city, searchTerm) {
			results = append(results, airports...)
		}
	}

	return results
}
