package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"u_demo/controller"
	ginzap "u_demo/gin_zap"
	"u_demo/middlewares"
)

func SetRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置为发布模式
	}
	r := gin.New()
	r.Use(ginzap.GinLogger(), ginzap.GinRecovery(true))

	v1 := r.Group("v1/api")

	v1.POST("/signup", controller.SignUpHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1.POST("/login", controller.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware()) // jwt 中间件
	{
		// 社区
		v1.GET("/community", controller.CommunityHandle)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		// 文章
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandlerByID)
		v1.GET("/posts", controller.GetPostListHandler)
		// 根据时间或分数获取帖子列表
		v1.GET("/post2", controller.GetPostListHandler2)
		//投票  用到redis 和 其中的算法
		v1.POST("/vote", controller.PostVoteController)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "404 NOT Found",
		})
	})

	return r
}
