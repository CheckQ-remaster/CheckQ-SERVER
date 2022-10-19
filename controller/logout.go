package controller

import (
	"net/http"

	"example.com/m/helper"
	"example.com/m/modules"
	"example.com/m/response"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {

	ad, err := helper.ExtractTokenMetadata(c)
	if err != nil {
		response.Response(c, http.StatusBadRequest, "log out token has not founded")
		return
	}

	deleted, err := modules.DeleteAuth(ad.AccessUuid)
	if err != nil || deleted == 0 {
		response.Response(c, http.StatusBadRequest, "log out delete auth error")
		return
	}

	response.Response(c, http.StatusAccepted, "log out suceess")
}
