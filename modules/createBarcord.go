package modules

import (
	"sync"

	"github.com/gin-gonic/gin"
)

func CreateBarCord(c *gin.Context, id string, seed string, wg sync.WaitGroup) {

	/*for {
		seedRand, _ := strconv.Atoi(seed)
		response.Response(c, http.StatusAccepted, seedRand)

		if c.Query("status") == "done" {

		}
	}

	wg.Done()*/
}
