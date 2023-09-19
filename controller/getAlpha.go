package controller

import (
	"net/http"

	"example.com/m/database"
	"example.com/m/models"
	"example.com/m/response"
	"github.com/gin-gonic/gin"
)

type alphaInfo struct {
	Alpha int
}

func GetAlpha(c *gin.Context) {

	hotelName := c.Query("hotel")
	roomNum := c.Query("room")

	alpha := &alphaInfo{}

	err := database.DB.Model(&models.Hotel{}).Select("alpha").Where("hotelname = ? and room = ?", hotelName, roomNum).Find(alpha).Error
	if err != nil {
		response.Response(c, http.StatusBadRequest, "get alpha select alpha erorr")
		return
	}

	response.Response(c, http.StatusAccepted, rand.Intn(alpha.Alpha))
}
