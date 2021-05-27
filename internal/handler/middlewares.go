package handler

import "github.com/gin-gonic/gin"

func enableCORS(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
	ctx.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
	ctx.Header("Content-Type", "application/json")

	ctx.Next()
}
