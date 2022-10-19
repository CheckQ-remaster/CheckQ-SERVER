package controller

import (
	"net/http"
	"strings"

	"example.com/m/database"
	"example.com/m/models"
	"example.com/m/response"
	"github.com/gin-gonic/gin"
)

func GetHotelInfo(c *gin.Context) {

	hotel := &models.Hotel{}

	hotelName := c.Query("hotel")
	hotelRoom := c.Query("room")

	if strings.Contains(hotelRoom, "호") {
		hotelRoom = strings.TrimRight(hotelRoom, "호")
	}

	err := database.DB.Model(&models.Hotel{}).Omit("alpha").Where("Hotelname = ? and Room = ?", hotelName, hotelRoom).Find(hotel).Error
	if err != nil {
		response.Response(c, http.StatusBadRequest, "get hotel info find error")
		return
	}

	response.Response(c, http.StatusAccepted, hotel)
}
