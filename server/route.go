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

		gd := api.Group("/get")
		{
			gd.GET("users-batch/:batchNumber", user.GetBachOfUSers)
			gd.GET("user-by-id/:id", user.GetUserByID)

			gd.GET("posts-batch/:batchNumber", post.GetBatchOfPosts)
			gd.GET("post-by-id/:id", post.GetPostByID)

			gd.GET("album-by-id/:id", albums.GetAlbumByID)

		}

		p := api.Group("/post")
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

		album := api.Group("/album")
		album.Use(mw.AuthMiddleware())
		{
			album.POST("create", albums.CreateAlbum)
			album.POST("update/:id", albums.UpdateAlbum)
			album.POST("delete", albums.DeleteAlbum)
		}

		d := api.Group("/data")
		d.Use(mw.AuthMiddleware())
		{
			d.POST("catch", data.SaveData)
		}
	}
}
