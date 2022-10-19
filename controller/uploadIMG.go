package controller

import (
	"net/http"

	"example.com/m/database"
	"example.com/m/modules"
	"example.com/m/response"
	"github.com/gin-gonic/gin"
)

func UploadIMG(c *gin.Context) {

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		response.Response(c, http.StatusBadRequest, "upload image formfile error")
		return
	}

	fileName, err := modules.ImageUpload(&file, header)
	if fileName == "" || err != nil {
		response.Response(c, http.StatusBadRequest, "upload image imageUpload error")
		return
	}

	hotelName := c.Query("hotel")

	if err = database.DB.Where("Hotel = ?", hotelName).Update("HotelImage = ?", fileName).Error; err != nil {
		response.Response(c, http.StatusBadRequest, "upload image create error")
		return
	}

	response.Response(c, http.StatusAccepted, "upload image create success")

}
