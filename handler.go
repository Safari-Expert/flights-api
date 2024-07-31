package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/krisukox/google-flights-api/flights"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

var session *flights.Session

func GetFlights(c *fiber.Ctx) error {
	offers, priceRange, err := session.GetOffers(
		context.Background(),
		flights.Args{
			Date:       time.Now().AddDate(0, 0, 30),
			ReturnDate: time.Now().AddDate(0, 0, 37),
			SrcCities:  []string{"Madrid"},
			DstCities:  []string{"Estocolmo"},
			Options: flights.Options{
				Travelers: flights.Travelers{Adults: 2},
				Currency:  currency.EUR,
				Stops:     flights.Stop1,
				Class:     flights.Economy,
				TripType:  flights.RoundTrip,
				Lang:      language.Spanish,
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	if priceRange != nil {
		fmt.Printf("High price %d\n", int(priceRange.High))
		fmt.Printf("Low price %d\n", int(priceRange.Low))
	}
	fmt.Println(offers)
	return c.SendString("All Flights")
}

func GetCity(c *fiber.Ctx) error {
	city := c.Query("city")
	if city == "" {
		return c.JSON(nil)
	}
	searchResult := searchCities(cityMap, city)
	return c.JSON(searchResult)
}
