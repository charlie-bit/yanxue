package api

import (
	"github.com/charlie-bit/yanxue/config"
	"github.com/charlie-bit/yanxue/pkg/constant"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Start() {
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(Cors())

	Register(g)

	go func() {
		err := g.Run(config.Cfg.Addr)
		if err != nil {
			panic(err)
		}
	}()
}

func Register(g *gin.Engine) {
	if config.Cfg.Env != constant.Prod {
		// nolint
		// path: http://127.0.0.1:8000/swagger/index.html#
		// swagger doc handler
		// doc := g.Group("/swagger", gin.BasicAuth(gin.Accounts{
		//	config.GCfg.Swagger.User: config.GCfg.Swagger.Password,
		// }))
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	v1 := g.Group("/api/v1")
	v1.GET(config.Cfg.HealthURI, func(context *gin.Context) {
		context.JSON(http.StatusOK, "hello world")
	})
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "*")
			c.Header("Access-Control-Allow-Headers", "*")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "")
			return
		}
		c.Next()
	}
}
