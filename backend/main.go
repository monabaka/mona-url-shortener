package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type application struct {
	DB     *gorm.DB
	Fiber  *fiber.App
	Domain string
	Port   string
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := ConnectToDB()
	db.AutoMigrate(&URL{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	app := &application{
		DB:     db,
		Domain: os.Getenv("DOMAIN"),
		Port:   os.Getenv("PORT"),
	}

	app.Fiber = fiber.New()

	app.Fiber.Post("/shorten", app.ShortenURL)
	app.Fiber.Get("/:shortID", app.RedirectURL)

	if app.Port == "" {
		app.Port = "8080"
	}

	log.Fatal(app.Fiber.Listen(":" + app.Port))
}
