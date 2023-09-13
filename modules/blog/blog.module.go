package blog

import (
	"google.golang.org/grpc"
	. "practice/auth/core/constants"
	"practice/auth/grpc/blog"
)

type BlogModule struct {
	Controller    *BlogController
	ServiceClient blog.BlogServiceClient
}

func InitBlogModule(GRPCMap map[string]*grpc.ClientConn, serviceMap map[string]string) *BlogModule {
	serviceAddress := serviceMap[BlogService]
	conn := GRPCMap[serviceAddress]
	serviceClient := blog.NewBlogServiceClient(conn)

	blogController := NewBlogController(serviceClient)

	return &BlogModule{Controller: blogController, ServiceClient: serviceClient}
}
