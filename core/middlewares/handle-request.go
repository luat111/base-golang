package middlewares

import (
	"net/http"
	"practice/auth/core/constants"
	. "practice/auth/core/interfaces"
	"reflect"

	"github.com/gin-gonic/gin"
)

func ValidateRequest[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := ResponseDefault{}
		req := new(T)
		statusCode := 0

		err := ctx.ShouldBind(req)

		if err != nil {
			res.Status = false
			res.Errors = []string{err.Error()}
			statusCode = http.StatusBadRequest
			ctx.AbortWithStatusJSON(statusCode, res)
			return
		}

		errs := ValidateStruct(req)
		if errs != nil {
			res.Status = false
			res.Errors = errs
			statusCode = http.StatusBadRequest
			ctx.AbortWithStatusJSON(statusCode, res)
			return
		}

		ctx.Set(constants.RequestPayload, req)
		ctx.Next()
	}
}

func ValidatePaginateRequest[T, R any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := PaginateResponse{}
		req := QueryRequest[T, R]{}
		err := ctx.ShouldBind(&req)
		statusCode := 0

		var queryFields = req.QueryFields
		var orderFields = req.OrderFields

		if err != nil {
			res.Status = false
			res.Errors = []string{err.Error()}
			statusCode = http.StatusBadRequest
			ctx.AbortWithStatusJSON(statusCode, res)
			return
		}

		errs := ValidateStruct(&queryFields)
		if errs != nil {
			res.Status = false
			res.Errors = errs
			statusCode = http.StatusBadRequest
			ctx.AbortWithStatusJSON(statusCode, res)
			return
		}

		var filter QueryRequest[T, R]
		if reflect.ValueOf(queryFields).IsZero() == false {
			filter.QueryFields = queryFields
		}

		if reflect.ValueOf(orderFields).IsZero() == false {
			filter.OrderFields = orderFields
		}

		ctx.Set(constants.RequestPayload, filter)
		return
	}
}
