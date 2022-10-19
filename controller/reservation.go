package controller

import (
	"net/http"
	"strings"

	"example.com/m/database"
	"example.com/m/models"
	"example.com/m/response"
	"github.com/gin-gonic/gin"
)

func Reservation(c *gin.Context) {

	hotel := &models.Hotel{}
	err := c.Bind(hotel)
	if err != nil {
		response.Response(c, http.StatusBadRequest, "reservation bind error")
	}

	if strings.Contains(hotel.Room, "호") {
		hotel.Room = strings.TrimRight(hotel.Room, "호")
	}

	id := c.Query("id")

	err = database.DB.Model(&models.Hotel{}).Where("hotelname = ? and room = ?", hotel.Hotelname, hotel.Room).Update("user_id", id).Error
	if err != nil {
		response.Response(c, http.StatusBadRequest, "reservation update error")
		return
	}

	response.Response(c, http.StatusAccepted, "reservation success")
}
