package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	e := echo.New()

	books = append(books, Book{ID: "1", Title: "Harry Potter dan Batu Bertuah", Author: "J.K. Rowling"})
	books = append(books, Book{ID: "2", Title: "Lord of the Rings", Author: "J.R.R. Tolkien"})

	e.GET("/books", getBooks)
	e.GET("/books/:id", getBook)
	e.POST("/books", createBook)
	e.PUT("/books/:id", updateBook)

	e.Start(":8000")
}

func getBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}

func getBook(c echo.Context) error {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			return c.JSON(http.StatusOK, book)
		}
	}

	return echo.NotFoundHandler(c)
}

func createBook(c echo.Context) error {
	book := new(Book)
	if err := c.Bind(book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	books = append(books, *book)

	return c.JSON(http.StatusCreated, book)
}

func updateBook(c echo.Context) error {
	id := c.Param("id")
	for index, book := range books {
		if book.ID == id {
			// Membuat objek book dengan data dari body permintaan
			updatedBook := &Book{
				ID:     id,
				Title:  book.Title,
				Author: book.Author,
			}
			if err := c.Bind(updatedBook); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			books[index] = *updatedBook
			return c.JSON(http.StatusOK, updatedBook)
		}
	}

	return echo.NotFoundHandler(c)
}
