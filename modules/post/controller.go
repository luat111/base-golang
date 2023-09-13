package post

import (
	"net/http"

	"practice/auth/core/constants"
	. "practice/auth/core/interfaces"
	"practice/auth/core/utils"
	. "practice/auth/modules/post/service"
	. "practice/auth/modules/post/model"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	Service *PostService
}

func NewPostController(postService *PostService) *PostController {
	return &PostController{Service: postService}
}

// @Tags		Posts
// @Security	BearerAuth
// @Accept		json
// @Produce	json
// @Success	200	{object}	PaginateResponse
// @Failure	500
// @Param		payload	body	CreatePostSchema	true	"c"
// @Router		/post/create [post]
func (pc *PostController) CreatePost(c *gin.Context) {
	var payload *CreatePostSchema
	res := ResponseDefault{}
	httpStatusCode := 0

	body, _ := c.Get(constants.RequestPayload)
	payload = body.(*CreatePostSchema)

	err, statusCode := pc.Service.CreatePost(payload)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		httpStatusCode = statusCode
	} else {
		res.Status = true
		httpStatusCode = http.StatusOK
	}

	c.JSON(httpStatusCode, res)
}

func (pc *PostController) UpdatePost(c *gin.Context) {
	var payload *UpdatePostSchema
	res := ResponseDefault{}
	id := c.Param("id")
	httpStatusCode := 0

	body, _ := c.Get(constants.RequestPayload)
	payload = body.(*UpdatePostSchema)

	err, statusCode := pc.Service.UpdatePost(id, payload)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		httpStatusCode = statusCode
	} else {
		res.Status = true
		httpStatusCode = http.StatusOK
	}

	c.JSON(httpStatusCode, res)
}

// @Security	BearerAuth
// @Tags		Posts
// @Accept		json
// @Produce	json
// @Success	200	{object}	PaginateResponse
// @Failure	500
// @Param		page	query	int												false	"a"	minimum(1)	default(1)
// @Param		size	query	int												false	"b"	default(10)
// @Param		payload	body	QueryRequest[QueryPostSchema, OrderPostSchema]	true	"c"
// @Router		/post/listing [post]
func (pc *PostController) GetListPost(c *gin.Context) {
	var filter QueryRequest[QueryPostSchema, OrderPostSchema]
	var size, page string
	res := PaginateResponse{}
	httpStatusCode := 0

	body, _ := c.Get(constants.RequestPayload)
	filter = body.(QueryRequest[QueryPostSchema, OrderPostSchema])

	size, page = c.Query("size"), c.Query("page")
	iPage, iSize := utils.PageSize(page, size)
	paginate := PaginateRequest{
		Page: iPage,
		Size: iSize,
	}

	data, err, statusCode := pc.Service.GetListPost(filter, paginate)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		httpStatusCode = statusCode
	} else {
		res.Status = true
		res.Data = &data
		httpStatusCode = http.StatusOK
	}

	c.JSON(httpStatusCode, res)
}

func (pc *PostController) GetDetailPost(c *gin.Context) {
	res := ResponseDefault{}
	id := c.Param("id")
	httpStatusCode := 0

	data, err, statusCode := pc.Service.GetDetailPost(id)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		httpStatusCode = statusCode
	} else {
		res.Status = true
		res.Data = &data
		httpStatusCode = http.StatusOK
	}

	c.JSON(httpStatusCode, res)
}

func (pc *PostController) DeletePost(c *gin.Context) {
	res := ResponseDefault{}
	id := c.Param("id")
	httpStatusCode := 0

	err, statusCode := pc.Service.DeletePost(id)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		httpStatusCode = statusCode
	} else {
		res.Status = true
		httpStatusCode = http.StatusOK
	}

	c.JSON(httpStatusCode, res)
}
