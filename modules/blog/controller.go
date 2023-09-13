package blog

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	. "practice/auth/core/interfaces"
	"practice/auth/grpc/blog"
	"time"
)

type BlogController struct {
	ServiceClient blog.BlogServiceClient
}

func NewBlogController(blogService blog.BlogServiceClient) *BlogController {
	return &BlogController{ServiceClient: blogService}
}

// @Tags		Blogs
// @Security BearerAuth
// @Accept		json
// @Produce	json
// @Success	200	{object}	PaginateResponse
// @Failure	500
// @Param		payload	body	CreateBlogSchema	true	"c"
// @Router		/blog/ [get]
func (bc *BlogController) GetBlog(c *gin.Context) {
	res := ResponseDefault{}
	httpStatusCode := 0

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	result, err := bc.ServiceClient.GetBlog(ctx, &blog.BlogRequest{
		Blog: &blog.Blog{
			Title: "Client test",
		},
	})

	if err != nil {
		res.Status = false
		res.Message = err.Error()
		httpStatusCode = http.StatusInternalServerError
	} else {
		res.Status = true
		res.Data = result
		httpStatusCode = http.StatusOK
		c.JSON(httpStatusCode, res)
	}
}
