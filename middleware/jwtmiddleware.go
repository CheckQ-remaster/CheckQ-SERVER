package middleware

import (
	"net/http"

	"example.com/m/helper"
	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"message": message,
	})
}

func TokenAuthMiddleware() gin.HandlerFunc { // token확인 middleware

	return func(c *gin.Context) {

		if err := helper.TokenVaild(c); err != nil {
			respondWithError(c, http.StatusUnauthorized, "vaild error")
		}

		ad, err := helper.ExtractTokenMetadata(c)
		if err != nil {
			respondWithError(c, http.StatusUnauthorized, "extract token error")
		}

		if _, err := helper.FetchAuth(ad); err != nil {
			respondWithError(c, http.StatusUnauthorized, "fetch auth error")
		}

		c.Next()
	}
}
