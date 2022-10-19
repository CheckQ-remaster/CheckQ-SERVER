package middleware

import (
	"net/http"

	"example.com/m/database"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

// 기업인지 일반 사용자인지 체크
func RoleMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		//c.Request.Header.Get("id")

		res := map[string]interface{}{}

		err := database.DB.Model(&models.User{}).Select("role").Where("id = ?", c.Request.Header.Get("id")).Find(&res).Error
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if res["role"] != "enterprise" {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			c.Next()
		}

	}
}
