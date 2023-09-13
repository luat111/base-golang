package user

import (
	"net/http"

	"practice/auth/core/constants"
	. "practice/auth/core/interfaces"
	"practice/auth/core/utils"
	. "practice/auth/modules/user/model"
	. "practice/auth/modules/user/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *UserService
}

func NewUserController(userService *UserService) *UserController {
	return &UserController{Service: userService}
}

// @Tags		Users
// @Security	BearerAuth
// @Accept		json
// @Produce	json
// @Success	200	{object}	PaginateResponse
// @Failure	500
// @Param		payload	body	CreateUserSchema	true	"c"
// @Router		/user/create [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var payload *CreateUserSchema
	res := ResponseDefault{}
	httpStatusCode := 0

	body, _ := c.Get(constants.RequestPayload)
	payload = body.(*CreateUserSchema)

	err, statusCode := uc.Service.CreateUser(payload)
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

func (uc *UserController) UpdateUser(c *gin.Context) {
	var payload *UpdateUserSchema
	res := ResponseDefault{}
	id := c.Param("id")
	httpStatusCode := 0

	body, _ := c.Get(constants.RequestPayload)
	payload = body.(*UpdateUserSchema)

	err, statusCode := uc.Service.UpdateUser(id, payload)
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
// @Tags		Users
// @Accept		json
// @Produce	json
// @Success	200	{object}	PaginateResponse
// @Failure	500
// @Param		page	query	int												false	"a"	minimum(1)	default(1)
// @Param		size	query	int												false	"b"	default(10)
// @Param		payload	body	QueryRequest[QueryUserSchema, OrderUserSchema]	true	"c"
// @Router		/user/listing [post]
func (uc *UserController) GetListUser(c *gin.Context) {
	var filter QueryRequest[QueryUserSchema, OrderUserSchema]
	var size, page string
	res := PaginateResponse{}
	httpStatusCode := 0

	body, _ := c.Get(constants.RequestPayload)
	filter = body.(QueryRequest[QueryUserSchema, OrderUserSchema])

	size, page = c.Query("size"), c.Query("page")
	iPage, iSize := utils.PageSize(page, size)
	paginate := PaginateRequest{
		Page: iPage,
		Size: iSize,
	}

	data, err, statusCode := uc.Service.GetListUser(filter, paginate)
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

func (uc *UserController) GetDetailUser(c *gin.Context) {
	res := ResponseDefault{}
	id := c.Param("id")
	httpStatusCode := 0

	data, err, statusCode := uc.Service.GetDetailUser(id)
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

func (uc *UserController) DeleteUser(c *gin.Context) {
	res := ResponseDefault{}
	id := c.Param("id")
	httpStatusCode := 0

	err, statusCode := uc.Service.DeleteUser(id)
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
