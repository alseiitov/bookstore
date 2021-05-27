package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookInput struct {
	Name   string `json:"name" binding:"required,max=255"`
	Author string `json:"author" binding:"required,max=255"`
}

func (h *Handler) GetAllBooks(ctx *gin.Context) {
	books, err := h.services.Books.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (h *Handler) GetBook(ctx *gin.Context) {
	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	book, err := h.services.Books.GetByID(bookID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (h *Handler) AddBook(ctx *gin.Context) {
	var input bookInput

	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Books.Create(input.Name, input.Author)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) UpdateBook(ctx *gin.Context) {
	var input bookInput

	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Books.Update(bookID, input.Name, input.Author)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) DeleteBook(ctx *gin.Context) {
	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Books.Delete(bookID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)
}
