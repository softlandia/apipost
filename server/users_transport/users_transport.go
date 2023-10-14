package users_transport

import (
	"apipost/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Transport struct {
	service.Service
}

func (t *Transport) List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, t.Repositories.Users.Select(""))
}

func (t *Transport) ById(ctx *gin.Context) {
	uid, err := strconv.Atoi(ctx.Param("uid"))
	if err != nil {
		ctx.JSON(400, gin.H{
			"error":   err.Error(),
			"message": "шарик ты балбес",
		})
		return
	}
	ctx.JSON(http.StatusOK, t.Repositories.Users.GET(uid))
}
