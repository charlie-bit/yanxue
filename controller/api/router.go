package api

import (
	"github.com/charlie-bit/yanxue/config"
	base "github.com/charlie-bit/yanxue/controller/common"
	"github.com/charlie-bit/yanxue/pkg/constant"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/charlie-bit/yanxue/docs" // swagger
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

	v1.GET("/phone_code", user.PhoneCode)
	v1.POST("/register", user.Register)
	v1.POST("/login", user.Login)
	v1.GET("/check_code/:phone", user.GetCheckCode)
	v1.POST("/verify_check_code", user.VerifyCheckCode)
	v1.GET("/sign_out", base.JwtAuthentication, user.SignOut)

	rc := g.Group(v1.BasePath() + "/role")
	rc.Use(base.JwtAuthentication)
	rc.POST("/new", role.New)
	rc.GET("/:id", role.GetByID)
	rc.DELETE("/:id", role.Del)
	rc.GET("/list", role.List)
	rc.PUT("/:id", role.Update)

	relc := g.Group(v1.BasePath() + "/relation")
	relc.Use(base.JwtAuthentication)
	relc.POST("/new", relation.Create)
	relc.GET("/:user_id", relation.GetUID)
	relc.GET("/list", relation.List)
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
