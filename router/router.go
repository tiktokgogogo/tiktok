package router

import (
	"github.com/gin-gonic/gin"
	"tiktok/api"
	"tiktok/config"
	"tiktok/middleware"
)

func InitRouter() {
	gin.SetMode(config.AppMode)
	r := gin.Default()
	apiRouter := r.Group("/douyin")
	apiRouter.GET("/feed/", middleware.JwtWithOutLogin(), api.Feed)
	apiRouter.POST("/publish/action/", middleware.JwtToken(), api.UpVideo)
	apiRouter.GET("/publish/list/", middleware.JwtToken(), api.VideoList)
	apiRouter.GET("/user/", middleware.JwtToken(), api.UserInfo)
	apiRouter.POST("/user/register/", api.Register)
	apiRouter.POST("/user/login/", api.Login)
	//互动
	apiRouter.POST("/favorite/action/", middleware.JwtToken(), api.Like)
	apiRouter.GET("/favorite/list/", middleware.JwtToken(), api.LikeList)
	apiRouter.POST("/comment/action/", middleware.JwtToken(), api.CommAction)
	apiRouter.GET("/comment/list/", middleware.JwtToken(), api.CommList)
	//社交
	apiRouter.POST("/relation/action/", middleware.JwtToken(), api.RelationAction)
	apiRouter.GET("/relation/follow/list/", middleware.JwtToken(), api.GetFollowing)
	apiRouter.GET("/relation/follower/list", middleware.JwtToken(), api.GetFollowers)
	apiRouter.GET("/relation/friend/list", middleware.JwtToken(), api.GetFileList)

	_ = r.Run(config.DBPort)
}
