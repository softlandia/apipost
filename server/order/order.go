package order

import (
	"apipost/model/order"
	"apipost/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handlers struct {
	service.Service
}

// Add - POST /order - создать заказ
func (t Handlers) Add(ctx *gin.Context) {
	ordr := order.Data{}
	if err := ctx.BindJSON(&ordr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := t.Repositories.Orders.Insert(ordr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"full_path": ctx.FullPath(),
		"method":    ctx.Request.Method,
	})
}

// Del - DEL  /order/{order_uid}
func (t Handlers) Del(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"full_path": ctx.FullPath(),
		"method":    ctx.Request.Method,
		"order_id":  ctx.Param("order_id"),
	})
}

// Get
// GET /order/{order_uid} - один заказ
func (t Handlers) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"full_path": ctx.FullPath(),
		"method":    ctx.Request.Method,
		"order":     t.Repositories.Orders.GET(ctx.Param("order_uid")),
	})
}

// List
// GET /orders - список заказов
func (t Handlers) List(ctx *gin.Context) {
	list := t.Repositories.Orders.Select(ctx.Query(""))
	ctx.JSON(http.StatusOK, gin.H{
		"full_path": ctx.FullPath(),
		"method":    ctx.Request.Method,
		"list":      list,
	})
}

// Update -
// PATCH /order/{order_uid}
func (t Handlers) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"full_path": ctx.FullPath(),
		"method":    ctx.Request.Method,
		"order_uid": ctx.Param("order_uid"),
	})
}
