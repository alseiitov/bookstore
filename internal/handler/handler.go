package handler

import (
	_ "github.com/alseiitov/bookstore/docs"
	"github.com/alseiitov/bookstore/internal/service"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()
	router.Use(enableCORS)

	api := router.Group("/api")
	{
		api.GET("/books", h.GetAllBooks)
		api.GET("/books/:id", h.GetBook)
		api.POST("/books", h.AddBook)
		api.PUT("/books/:id", h.UpdateBook)
		api.DELETE("/books/:id", h.DeleteBook)
	}

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
