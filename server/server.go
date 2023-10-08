package server

import (
	"apipost/server/order"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	Port string
}

func New(port string) *gin.Engine {
	g := gin.Default()
	orders := order.Handlers{}

	g.GET("/about", func(ctx *gin.Context) {
		ctx.JSON(http.StatusNoContent, nil)
	})

	g.GET("/orders", orders.List)
	g.GET("/order/:order_id", orders.Get)
	g.POST("/order", orders.Add)
	g.PATCH("/order/:order_id", orders.Update)
	g.DELETE("/order/:order_id", orders.Del)

	return g
}
