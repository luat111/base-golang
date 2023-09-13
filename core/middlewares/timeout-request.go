package middlewares

import (
	"context"
	"net/http"
	"time"

	. "practice/auth/core/interfaces"

	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(perSecond int) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := ResponseDefault{Status: false}

		convertToSecond := time.Duration(perSecond) * time.Second

		ctx, cancel := context.WithTimeout(c.Request.Context(), convertToSecond)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		done := make(chan bool, 1)
		go func() {
			c.Next()
			done <- true
		}()

		select {
		case <-done:
			return
		case <-ctx.Done():
			res.Message = "Request timeout"
			c.AbortWithStatusJSON(http.StatusRequestTimeout, res)
		}
	}
}
