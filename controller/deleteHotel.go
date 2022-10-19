package controller

import (
	"net/http"

	"example.com/m/database"
	"example.com/m/models"
	"example.com/m/response"
	"github.com/gin-gonic/gin"
)

func DeleteHotel(c *gin.Context) {

	hotel := &models.Hotel{}

	err := database.DB.Model(&models.Hotel{}).Delete(hotel).Error
	if err != nil {
		response.Response(c, http.StatusBadRequest, "delete hotel delete error")
		return
	}

	response.Response(c, http.StatusAccepted, "delete hotel success")
}
