package main

import (
	"api/config"
	"api/db"
	"api/infra/api/handler"
	"api/infra/persistance"
	"api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.GetConfig()
	db := db.NewDB(config)

	postRepository := persistance.NewPostRepository(db)
    postUsecase := usecase.NewPostUsecase(postRepository)
    postHandler := handler.NewPostHandler(postUsecase)


	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": db,
		})
	})

	postsGroup := r.Group("/posts")
	postsGroup.GET("", postHandler.GetPosts)
	postsGroup.POST("", postHandler.CreatePost)
	postsGroup.GET(":id", postHandler.GetPost)
	// postsGroup.PUT(":id", postHandler.UpdatePost)
	postsGroup.DELETE(":id", postHandler.DeletePost)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
