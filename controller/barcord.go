package controller

import (
	"github.com/gin-gonic/gin"
)

var seedRand string

func Barcord(c *gin.Context) {

	/*id := c.Query("id")

	rand := c.Query("rand")

	temp := &models.User{}
	//barUser := &models.BarcordUser{}

	if id != "" {

		barUser.IdUser = id
		barUser.Status = "run"

		if err := database.DB.Create(barUser).Error; err != nil {
			response.Response(c, http.StatusBadRequest, "barcord cerate error")
			return
		}

		if err := database.DB.Where("id = ?", id).Find(temp).Error; err != nil {
			response.Response(c, http.StatusBadRequest, "barcord find error")
			return
		}

		seedRand := temp.Seed
		response.Response(c, http.StatusAccepted, seedRand)

	} else {
		if rand == seedRand {
			response.Response(c, http.StatusAccepted, "OK")
		} else {
			response.Response(c, http.StatusAccepted, "NO")
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go modules.CreateBarCord(c, id, temp.Seed, wg)
	wg.Wait()*/
}
