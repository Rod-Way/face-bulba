package server

import (
	albums "faceBulba/internal/albums"
	data "faceBulba/internal/data"
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

		api.POST("catch-data", data.SaveData)

		p := api.Group("/posts")
		p.Use(mw.AuthMiddleware())
		{
			p.POST("create", post.CreatePost)
			p.POST("update/:id", post.UpdatePost)
			p.POST("delete", post.DeletePost)
		}

		comm := api.Group("/comment")
		comm.Use(mw.AuthMiddleware())
		{
			comm.POST("create", post.CreateComment)
			comm.POST("update/:id", post.UpdateComment)
			comm.POST("delete", post.DeleteComment)
		}

		album := api.Group("/albums")
		album.Use(mw.AuthMiddleware())
		{
			album.POST("create", albums.CreateAlbum)
		}

	}
}
