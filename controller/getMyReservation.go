package controller

import (
	"net/http"

	"example.com/m/database"
	"example.com/m/models"
	"example.com/m/response"
	"github.com/gin-gonic/gin"
)

func GetMyReservation(c *gin.Context) {

	id := c.Query("id")

	hotel := []models.Hotel{}

	// err := database.DB.Select("id").Model(&models.User{}).Preload("Hotels").Where("id = ?", id).Find(&dbUser).Error // preload - 연관테이블 같이 불러오기
	// if err != nil {
	// 	response.Response(c, http.StatusBadRequest, "get my reservation find errpr")
	// 	return
	// }

	if err := database.DB.Where("user_id = ?", id).Find(&hotel).Error; err != nil {
		response.Response(c, http.StatusAccepted, "get my reservation error")
		return
	}

	response.Response(c, http.StatusAccepted, hotel)
}
