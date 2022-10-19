package controller

import (
	"net/http"

	"example.com/m/database"
	"example.com/m/helper"
	"example.com/m/models"
	"example.com/m/response"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	user := &models.User{}

	if err := c.Bind(user); err != nil {
		response.Response(c, http.StatusBadRequest, "login bind error")
		return
	}

	userPw := user.Pw

	findRes := database.DB.Where("id = ?", user.Id).Find(user)

	if findRes.RowsAffected == 0 {
		response.Response(c, http.StatusBadRequest, "Id had not found")
		return
	}

	res := helper.CheckPasswordHash(user.Pw, userPw)
	if !res {
		response.Response(c, http.StatusBadRequest, "Pw is not true")
		return
	}

	tk, err := helper.CreateJWT(user.Id) // 유저 id로 토큰 제작 - acceess, refresh
	if err != nil {
		response.Response(c, http.StatusBadRequest, "login create token error")
		return
	}

	err = helper.CreateAuth(tk) // 만들어진 토큰을 redis에 저장
	if err != nil {
		response.Response(c, http.StatusBadRequest, "login create auth error")
		return
	}

	c.SetCookie("access-Token", tk.AccessToken, int(tk.AtExpires), "/", "localhost", false, false)
	//c.SetCookie("refresh-Token", tk.RefreshToken, int(tk.RtExpires), "/", "localhost", false, false)

	/*token := map[string]string{
		"access_token": tk.AccessToken,
		//"refresh_token": tk.RefreshToken,
	}*/

	response.ResponseToken(c, http.StatusAccepted, tk.AccessToken, user.Id)
}
