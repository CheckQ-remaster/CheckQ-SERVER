package controller

import (
	"net/http"

	"example.com/m/database"
	"example.com/m/models"
	"example.com/m/response"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	id := c.Query("id")

	user := &models.User{}

	err := c.Bind(user)
	if err != nil {
		response.Response(c, http.StatusBadRequest, "get user bind error")
		return
	}

	if err = database.DB.Omit("pw").Where("id = ?", id).Find(user).Error; err != nil {
		response.Response(c, http.StatusBadRequest, "get user find error")
		return
	}

	response.Response(c, http.StatusAccepted, user)
}
