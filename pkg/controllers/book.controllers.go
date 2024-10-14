package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/book-crud-golang/db"
	"github.com/sherwin-77/book-crud-golang/pkg/models"
)

type BookRequest struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}

func CreateBook(ctx echo.Context) error {
	DB := db.DB

	bookRequest := BookRequest{}

	if err := ctx.Bind(&bookRequest); err != nil {
		return err
	}

	if err := ctx.Validate(bookRequest); err != nil {
		return err
	}

	book := models.Book{
		Title:  bookRequest.Title,
		Author: bookRequest.Author,
	}

	if err := DB.Create(&book).Error; err != nil {
		return err
	}

	return ctx.JSON(201, book)
}

func GetBooks(ctx echo.Context) error {
	DB := db.DB

	books := []models.Book{}

	if err := DB.Find(&books).Error; err != nil {
		return err
	}

	return ctx.JSON(200, books)
}

func GetBookByID(ctx echo.Context) error {
	DB := db.DB

	book := models.Book{}

	if err := DB.First(&book, "id = ?", ctx.Param("id")).Error; err != nil {
		return err
	}

	return ctx.JSON(200, book)
}

func UpdateBook(ctx echo.Context) error {
	DB := db.DB

	book := models.Book{}

	if err := DB.First(&book, "id = ?", ctx.Param("id")).Error; err != nil {
		return err
	}

	bookRequest := BookRequest{}

	if err := ctx.Bind(&bookRequest); err != nil {
		return err
	}

	if err := ctx.Validate(bookRequest); err != nil {
		return err
	}

	book.Title = bookRequest.Title
	book.Author = bookRequest.Author

	if err := DB.Save(&book).Error; err != nil {
		return err
	}

	return ctx.JSON(200, book)
}

func DeleteBook(ctx echo.Context) error {
	DB := db.DB

	book := models.Book{}

	if err := DB.First(&book, "id = ?", ctx.Param("id")).Error; err != nil {
		return err
	}

	if err := DB.Delete(&book).Error; err != nil {
		return err
	}

	return ctx.NoContent(204)
}
