package response

import (
	"log"

	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, code int, ms interface{}) {

	if v, ok := ms.(string); ok {
		ifString(c, code, v)
	} else if v, ok := ms.([]string); ok {
		ifStringSlice(c, code, v)
	} else if v, ok := ms.([]models.Hotel); ok {
		ifHotelSlice(c, code, v)
	} else if v, ok := ms.([]models.User); ok {
		ifUserSlice(c, code, v)
	} else if v, ok := ms.(*models.User); ok {
		ifUser(c, code, v)
	} else if v, ok := ms.(*models.Hotel); ok {
		ifHotel(c, code, v)
	} else if v, ok := ms.(int); ok {
		ifInt(c, code, v)
	} else {
		log.Println("response 무슨 타입인지 모르겠음")
	}
}

func ifString(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"message": msg,
	})
}

func ifStringSlice(c *gin.Context, code int, msg []string) {
	c.JSON(code, gin.H{
		"data": msg,
	})
}

func ifUserSlice(c *gin.Context, code int, msg []models.User) {
	c.JSON(code, gin.H{
		"data": msg,
	})
}

func ifHotelSlice(c *gin.Context, code int, msg []models.Hotel) {
	c.JSON(code, gin.H{
		"data": msg,
	})
}

func ResponseToken(c *gin.Context, code int, token string, id string) {
	c.JSON(code, gin.H{
		"token": token,
		"id":    id,
	})
}

func ifUser(c *gin.Context, code int, msg *models.User) {
	c.JSON(code, gin.H{
		"data": msg,
	})
}
func ifHotel(c *gin.Context, code int, msg *models.Hotel) {
	c.JSON(code, gin.H{
		"data": msg,
	})
}

func ifInt(c *gin.Context, code int, msg int) {
	c.JSON(code, gin.H{
		"qr": msg,
	})
}
