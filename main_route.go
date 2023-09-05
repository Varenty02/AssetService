package main

import (
	appctx "AssetService/component"
	"AssetService/middleware"
	"AssetService/response/transport"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoute(ctx appctx.AppContext, v1 *gin.RouterGroup) {
	v1 = v1.Group("/", middleware.Recover(ctx))
	//test
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	v1.GET("/ipinfo", transport.ListIps(ctx))

}
