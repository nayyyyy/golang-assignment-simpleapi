package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID      string `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"desc"`
}

var Books = []Book{}

func IndexBook(ctx *gin.Context) {
	if len(Books) == 0 {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status": "No content",
			"message": "Data not found",
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Data buku tersedia",
		"data":    Books,
	})
}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = fmt.Sprintf("%d", len(Books)+1)

	Books = append(Books, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "OK",
		"message": "Buku berhasil disimpan",
	})
}

func ShowBook(ctx *gin.Context) {
	bookID := ctx.Param("id")
	condition := false

	var showBook Book

	for i, book := range Books {
		if bookID == book.BookID {
			condition = true
			showBook = Books[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Data tidak ditemukan",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Data berhasil ditemukan",
		"data":    showBook,
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("id")
	condition := false

	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	for i, book := range Books {
		if bookID == book.BookID {
			condition = true
			Books[i] = updatedBook
			Books[i].BookID = bookID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Data tidak ditemukan",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Data berhasil diperbaharui",
		"data":    updatedBook,
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("id")
	condition := false

	var bookIndex int

	for i, book := range Books {
		if bookID == book.BookID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Data tidak ditemukan",
		})
		return
	}

	copy(Books[bookIndex:], Books[bookIndex+1:])
	Books[len(Books)-1] = Book{}
	Books = Books[:len(Books)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"id":      bookIndex + 1,
		"message": "Data berhasil dihapus",
	})
}
