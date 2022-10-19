package main

import (
	"example.com/m/database"
	"example.com/m/redis"
	"example.com/m/route"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default() // gin engine

	redis.InitRedis()  // redis 연결
	database.Connect() // mysql 연결

	route.Route(r) // 라우팅

	defer redis.CloseRedis() // redis 닫기

	r.Run() // 서버 실행
}
