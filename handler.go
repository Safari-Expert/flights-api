package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/krisukox/google-flights-api/flights"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

var session *flights.Session

func GetFlights(c *fiber.Ctx) error {
	var flightRequest FlightRequest
	if err := c.BodyParser(&flightRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	offers, priceRange, err := session.GetOffers(
		context.Background(),
		flights.Args{
			Date:       flightRequest.Date,
			ReturnDate: flightRequest.ReturnDate,
			SrcCities:  []string{flightRequest.SrcCity},
			DstCities:  []string{flightRequest.DstCity},
			Options: flights.Options{
				Travelers: flights.Travelers{
					Adults:       flightRequest.Adults,
					Children:     flightRequest.Children,
					InfantInSeat: flightRequest.InfantInSeat,
					InfantOnLap:  flightRequest.InfantOnLap,
				},
				Currency: currency.USD,
				Stops:    flights.Stops(flightRequest.Stops),
				Class:    flights.Class(flights.Economy),
				TripType: flights.OneWay,
				Lang:     language.AmericanEnglish,
			},
		},
	)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"offers":     offers,
		"priceRange": priceRange,
	})
}

func GetCity(c *fiber.Ctx) error {
	city := c.Query("city")
	if city == "" {
		return c.JSON(nil)
	}
	searchResult := searchCities(cityMap, city)
	return c.JSON(searchResult)
}
