package server

import (
	"apipost/server/order"
	"apipost/server/users_transport"
	"apipost/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	Port string
}

func New(port string, srvc *service.Service) *gin.Engine {
	g := gin.Default()

	g.GET("/about", func(ctx *gin.Context) {
		ctx.JSON(http.StatusNoContent, nil)
	})

	// работа с пользователями
	{
		users := users_transport.Transport{*srvc}
		g.GET("/users/:uid", users.ById)
		g.GET("/users", users.List)
	}

	// работа с заказами
	{
		orders := order.Handlers{*srvc}
		g.GET("/orders", orders.List)
		g.GET("/order/:order_uid", orders.Get)
		g.POST("/order", orders.Add)
		g.PATCH("/order/:order_uid", orders.Update)
		g.DELETE("/order/:order_uid", orders.Del)
	}

	return g
}
