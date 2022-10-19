package controller

import (
	"net/http"

	"example.com/m/database"
	"example.com/m/models"
	"example.com/m/response"
	"github.com/gin-gonic/gin"
)

func GetHotels(c *gin.Context) {

	hotel := []models.Hotel{}

	err := database.DB.Raw("select *from hotels order by rand() limit 20").Scan(&hotel).Error
	if err != nil {
		response.Response(c, http.StatusBadRequest, "get hotels build raw error")
		return
	}

	response.Response(c, http.StatusAccepted, hotel)

}
