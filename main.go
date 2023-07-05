package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rizkyrsyd28/recloth-backend/internal/delivery"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal memuat file .env")
	}

	f := fiber.New()

	delivery.MainRoutes(f)

	log.Fatal(f.Listen(":8080"))
}
