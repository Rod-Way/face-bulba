package server

import (
	mw "faceBulba/internal/middlewares"
	post "faceBulba/internal/posts"
	user "faceBulba/internal/user"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	api := route.Group("/api")
	{
		api.POST("register", user.AddUser)
		api.POST("login", user.LoginUser)
		api.GET("is-auth/:token", user.CheckUser)

		api.GET("get-batch/:batchNumber", post.GetBatchOfPosts)
		api.GET("get-by-id/:postID", post.GetPostByID)

		p := api.Group("/posts")
		p.Use(mw.AuthMiddleware())
		{
			p.POST("create", post.CreatePost)
			p.POST("update", post.UpdatePost)
			p.POST("delete", post.DeletePost)
		}
	}
}
