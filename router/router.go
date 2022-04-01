package router

import (
	"apiserver/handler/user"
	"net/http"

	"apiserver/handler/sd"
	"apiserver/router/middleware"

	"github.com/gin-gonic/gin"
)

func Land(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middleware
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// The health check handler
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	u := g.Group("users")
	{
		u.POST("", user.Create) //创建用户
	}

	return g
}
