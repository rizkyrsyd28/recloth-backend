package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	f.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Credentials", "true")
		return c.Next()
	})

	corsConf := cors.Config{
		AllowOrigins: "http://localhost:5137",
		//AllowMethods:     "GET, POST, HEAD, OPTIONS, PUT, DELETE, PATCH",
		//AllowHeaders:     "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With",
		//ExposeHeaders:    "Origin",
		AllowCredentials: true,
	}

	f.Use(cors.New(corsConf))

	delivery.MainRoutes(f)

	log.Fatal(f.Listen(":8080"))
}
