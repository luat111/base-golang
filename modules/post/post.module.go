package post

import (
	postService "practice/auth/modules/post/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type PostModule struct {
	Controller *PostController
	Service    *postService.PostService
}

func InitPostModule(NoSqlDb *mongo.Database) *PostModule {
	postService := postService.NewPostService(NoSqlDb)
	postController := NewPostController(postService)

	return &PostModule{Controller: postController, Service: postService}
}
