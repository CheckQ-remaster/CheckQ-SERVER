package controller

import (
	"net/http"

	"example.com/m/database"
	"example.com/m/models"
	"example.com/m/response"
	"github.com/gin-gonic/gin"
)

func GetHotel(c *gin.Context) {

	hotelName := c.Query("hotel")

	hotel := &models.Hotel{}

	err := database.DB.Model(&models.Hotel{}).Omit("Alpha").Where("Hotelname = ?", hotelName).Find(hotel).Error
	if err != nil {
		response.Response(c, http.StatusBadRequest, "get hotel  error")
		return
	}

	response.Response(c, http.StatusAccepted, hotel)
}
