package book

import (
	"fib/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func Getbooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		return c.Status(500).SendString("Error retrieving books")
	}
	return c.JSON(books)
}

func Getbook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		return c.Status(404).SendString("Book not found")
	}
	return c.JSON(book)
}

func Newbook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(Book)

	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).SendString("Invalid input")
	}

	if book.Title == "" || book.Author == "" || book.Rating == 0 {
		return c.Status(400).SendString("Title, Author, and Rating are required fields")
	}

	db.Create(&book)
	return c.JSON(book)
}

func Deletebook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		return c.Status(404).SendString("No book found with given ID")
	}

	db.Delete(&book)
	return c.SendString("Book successfully deleted")
}
