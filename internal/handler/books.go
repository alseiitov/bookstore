package handler

import (
	"net/http"
	"strconv"

	_ "github.com/alseiitov/bookstore/internal/model"
	"github.com/alseiitov/bookstore/internal/service"
	"github.com/gin-gonic/gin"
)

type bookInput struct {
	Name   string `json:"name" binding:"required,max=255"`
	Author string `json:"author" binding:"required,max=255"`
}

type errorResponse struct {
	Error string `json:"error"`
}

// @Summary Get all books
// @Tags books
// @ModuleID GetAllBooks
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Book
// @Failure default {object} errorResponse
// @Router /books [get]
func (h *Handler) GetAllBooks(ctx *gin.Context) {
	books, err := h.services.Books.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

// @Summary Get book
// @Tags books
// @ModuleID GetBook
// @Accept  json
// @Produce  json
// @Param id path int true "ID of book"
// @Success 200 {object} model.Book
// @Failure default {object} errorResponse
// @Router /books/{id} [get]
func (h *Handler) GetBook(ctx *gin.Context) {
	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	book, err := h.services.Books.GetByID(bookID)
	if err != nil {
		if err == service.ErrBookNotFound {
			ctx.JSON(http.StatusNotFound, errorResponse{err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse{err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, book)
}

// @Summary Add book
// @Tags books
// @ModuleID AddBook
// @Accept  json
// @Produce  json
// @Param input body bookInput true "book info"
// @Success 201 "ok"
// @Failure default {object} errorResponse
// @Router /books [post]
func (h *Handler) AddBook(ctx *gin.Context) {
	var input bookInput

	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	err = h.services.Books.Create(input.Name, input.Author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

// @Summary Update book
// @Tags books
// @ModuleID UpdateBook
// @Accept  json
// @Produce  json
// @Param id path int true "ID of book"
// @Param input body bookInput true "book info"
// @Success 200 "ok"
// @Failure default {object} errorResponse
// @Router /books/{id} [put]
func (h *Handler) UpdateBook(ctx *gin.Context) {
	var input bookInput

	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	err = h.services.Books.Update(bookID, input.Name, input.Author)
	if err != nil {
		if err == service.ErrBookNotFound {
			ctx.JSON(http.StatusNotFound, errorResponse{err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse{err.Error()})
		}
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary Delete book
// @Tags books
// @ModuleID DeleteBook
// @Accept  json
// @Produce  json
// @Param id path int true "ID of book"
// @Success 204 "ok"
// @Failure default {object} errorResponse
// @Router /books/{id} [delete]
func (h *Handler) DeleteBook(ctx *gin.Context) {
	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	err = h.services.Books.Delete(bookID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
