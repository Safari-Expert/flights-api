package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/krisukox/google-flights-api/flights"
)

func main() {
	var err error
	session, err = flights.New()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	log.Println("Session created")
	cityMap, err = createCityMap()
	if err != nil {
		log.Fatalf("Failed to create city map: %v", err)
	}
	log.Println("City map created")
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	app.Post("/flights", GetFlights)
	app.Get("/city", GetCity)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	app.Listen(":" + port)
}
