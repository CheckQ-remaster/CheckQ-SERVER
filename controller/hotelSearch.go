package controller

import (
	"net/http"

	"example.com/m/database"
	"example.com/m/models"
	"example.com/m/response"
	"github.com/gin-gonic/gin"
)

func HotelSearch(c *gin.Context) {

	hotelName := c.Query("hotel")

	hotel := []models.Hotel{}

	err := database.DB.Raw("select * from hotels where hotelname like ?", "%"+hotelName+"%").Scan(&hotel).Error
	if err != nil {
		response.Response(c, http.StatusBadRequest, "hotel search scan error")
		return
	}

	response.Response(c, http.StatusAccepted, hotel)
}
