package controller

import (
	"net/http"

	"example.com/m/database"
	"example.com/m/helper"
	"example.com/m/models"
	"example.com/m/response"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {

	user := &models.User{}

	err := c.Bind(user) // bind
	if err != nil {
		response.Response(c, http.StatusBadRequest, "signup bind err")
		return
	}

	if err := database.DB.Find("id = ?", user.Id).Error; err == nil { // id존재 여부 확인
		response.Response(c, http.StatusBadRequest, "signup id already exist")
		return
	}

	user.Pw, err = helper.HashPassword(user.Pw) // pw 암호화
	if err != nil {
		response.Response(c, http.StatusBadRequest, "signup hash pwd error")
		return
	}

	if user.Role != "" {
		user.Role = "enterprise" // 기업
	}

	if err := database.DB.Create(user).Error; err != nil { // user정보 insert
		response.Response(c, http.StatusBadRequest, "signup create row error")
		return
	}

	response.Response(c, http.StatusAccepted, "signup success")
}
