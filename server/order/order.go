package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handlers struct {
}

// Add - POST /order - создать заказ
func (o Handlers) Add(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"full_path": ctx.FullPath(),
		"method":    ctx.Request.Method,
	})
}

// Del - DEL  /order/{order_uid}
func (o Handlers) Del(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"full_path": ctx.FullPath(),
		"method":    ctx.Request.Method,
		"order_id":  ctx.Param("order_id"),
	})
}

// Get
// GET /order/{order_uid} - один заказ
func (o Handlers) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"full_path": ctx.FullPath(),
		"method":    ctx.Request.Method,
		"order_id":  ctx.Param("order_id"),
	})
}

// List
// GET /orders - список заказов
func (o Handlers) List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"full_path": ctx.FullPath(),
		"method":    ctx.Request.Method,
	})
}

// Update -
// PATCH /order/{order_uid}
func (o Handlers) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"full_path": ctx.FullPath(),
		"method":    ctx.Request.Method,
		"order_id":  ctx.Param("order_id"),
	})
}
