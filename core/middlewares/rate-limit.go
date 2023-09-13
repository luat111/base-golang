package middlewares

import (
	"fmt"
	"net/http"
	"time"

	. "practice/auth/core/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RateLimitMiddleware(redisClient *redis.Client, maxRequests int, perSecond int) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := ResponseDefault{Status: false}

		clientIP := c.ClientIP()

		rateLimitKey := fmt.Sprintf("rate-limit:%s", clientIP)

		sendRequestAndHandleError := func() (int64, error) {
			count, err := redisClient.Incr(c, rateLimitKey).Result()
			if err != nil {
				return 0, err
			}

			if count == 1 {
				convertToSecond := time.Duration(perSecond) * time.Second
				redisClient.Expire(c, rateLimitKey, convertToSecond)
			}

			return count, nil
		}

		count, err := sendRequestAndHandleError()
		if err != nil {
			res.Message = err.Error()

			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}

		if count > int64(maxRequests) {
			res.Message = "Too Many Request"

			c.AbortWithStatusJSON(http.StatusTooManyRequests, res)
			return
		}

		c.Next()
	}
}
