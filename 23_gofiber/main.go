package main

import (
	"fib/book"
	"fib/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func homepage(c *fiber.Ctx) error {
	return c.SendString("Hello world")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", homepage)
	app.Get("/api/v1/book", book.Getbooks)
	app.Get("/api/v1/book/:id", book.Getbook)
	app.Post("/api/v1/book/", book.Newbook)
	app.Delete("/api/v1/book/:id", book.Deletebook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	fmt.Println("Go fiber tutorial")

	app := fiber.New()
	initDatabase()

	// Ensure proper closure of the database connection
	db, err := database.DBConn.DB()
	if err != nil {
		fmt.Println("Error getting DB instance:", err)
	} else {
		defer db.Close()
	}

	setupRoutes(app)

	err = app.Listen(":3000")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
