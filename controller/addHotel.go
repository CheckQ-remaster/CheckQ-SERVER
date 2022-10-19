package controller

import (
	"math/rand"
	"net/http"

	"example.com/m/database"
	"example.com/m/models"
	"example.com/m/response"
	"github.com/gin-gonic/gin"
)

func AddHotel(c *gin.Context) {

	hotel := &models.Hotel{}

	err := c.Bind(hotel)
	if err != nil {
		response.Response(c, http.StatusBadRequest, "add hotel bind error")
		return
	}

	hotel.Alpha = rand.Intn(999999) + 9999

	err = database.DB.Omit("UserId").Create(hotel).Error
	if err != nil {
		response.Response(c, http.StatusBadRequest, "add hotel create error")
		return
	}

	response.Response(c, http.StatusAccepted, "add hotel success")
}
